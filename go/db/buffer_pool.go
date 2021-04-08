package db

import (
	"errors"
	"fmt"
)

// data store in databases named "pages", one page size like 16KB in mysql, 8KB in PostgreSQL.
// pages need to be cached in memory to accelerate access speed.

// page in memory use page table indicate.

const MaxPoolSize = 5

type BufferPoolManager struct {
	diskManager DiskManager        // manage page store in disk and read in memory
	pages       [MaxPoolSize]*Page // all pages
	replacer    *ClockReplacer     // Algorithm replace unused page
	freelist    []FrameID          // free space of frame, and unused page
	pageTable   map[PageID]FrameID // page indicate in memory
}

func NewBufferPoolManager(diskManager DiskManager, clockReplacer *ClockReplacer) *BufferPoolManager {
	freelist := make([]FrameID, 0)
	pages := [MaxPoolSize]*Page{}
	for i := 0; i < MaxPoolSize; i++ {
		freelist = append(freelist, FrameID(i))
		pages[FrameID(i)] = nil
	}

	return &BufferPoolManager{
		diskManager: diskManager,
		pages:       pages,
		replacer:    clockReplacer,
		freelist:    freelist,
		pageTable:   make(map[PageID]FrameID),
	}
}

func (b *BufferPoolManager) NewPage() *Page {

}
func (b *BufferPoolManager) FetchPage(pageID PageID) *Page {

}
func (b *BufferPoolManager) FlushPage(pageID PageID) bool {

}
func (b *BufferPoolManager) FlushAllPages() {

}
func (b *BufferPoolManager) DeletePage(pageID PageID) error {

}
func (b *BufferPoolManager) UnpinPage(pageID PageID, isDirty bool) error {

}

type PageID int

const pageSize = 5

// Page represent data structure in database
type Page struct {
	id       PageID
	pinCount int            // trace the number of concurrent access this page, if > 0, page be pinned
	isDirty  bool           // id dirty page: page modified after read from disk
	data     [pageSize]byte // page data
}

// DiskManager manage page store and read with disk, allocate etc.
type DiskManager interface {
	ReadPage(id PageID) (*Page, error) // read page from disk
	WritePage(page *Page) error        // write page to disk
	AllocatePage() *PageID             // allocate a new page
	DeallocatePage(id PageID)          // remove a page
}

// DiskManagerMock indicate the DiskManager interface Implement
type DiskManagerMock struct {
	numPage int // tract the number of page
	pages   map[PageID]*Page
}

func NewDiskManagerMock() *DiskManagerMock {
	return &DiskManagerMock{
		numPage: -1,
		pages:   make(map[PageID]*Page),
	}
}

// ReadPage read page from disk
func (d *DiskManagerMock) ReadPage(id PageID) (*Page, error) {
	if page, ok := d.pages[id]; ok {
		return page, nil
	}
	return nil, errors.New("page not found")
}

// WritePage write page to disk
func (d *DiskManagerMock) WritePage(page *Page) error {
	d.pages[page.id] = page
	return nil
}

// AllocatePage allocate a new page
func (d *DiskManagerMock) AllocatePage() *PageID {
	if d.numPage == DiskMaxNumPages-1 {
		return nil
	}
	d.numPage = d.numPage + 1
	pageID := PageID(d.numPage)
	return &pageID
}

// DeallocatePage delete page
func (d *DiskManagerMock) DeallocatePage(id PageID) {
	delete(d.pages, id)
}

// replacement algorithm
// such like LRU、FIFO、LFU

type FrameID int

type ClockReplacer struct {
	cList     *circularList
	clockHand **node
}

// Pin pin a frame, indicate this page maybe used later
func (c *ClockReplacer) Pin(id FrameID) {
	node := c.cList.find(id)
	if node == nil {
		return
	}

	if (*c.clockHand) == node {
		c.clockHand = &(*c.clockHand).next
	}

	c.cList.remove(id)
}

// Unpin indicate page unused
func (c *ClockReplacer) UnPin(id FrameID) {
	if !c.cList.hasKey(id) {
		c.cList.insert(id, true)
		if c.cList.size == 1 {
			c.clockHand = &c.cList.head
		}
	}
}

// Victim removes the victim frame as defined by the replacement policy
func (c *ClockReplacer) Victim() *FrameID {
	if c.cList.size == 0 {
		return nil
	}

	var victimFrameID *FrameID
	currentNode := (*c.clockHand)
	for {

		if currentNode.value.(bool) {
			currentNode.value = false
			c.clockHand = &currentNode.next
		} else {
			frameID := currentNode.key.(FrameID)
			victimFrameID = &frameID

			c.clockHand = &currentNode.next

			c.cList.remove(currentNode.key)
			return victimFrameID
		}
	}
}

// Size returns the size of the clock
func (c *ClockReplacer) Size() int {
	return c.cList.size
}

type node struct {
	key   interface{}
	value interface{}
	next  *node
	prev  *node
}

type circularList struct {
	head     *node
	tail     *node
	size     int
	capacity int
}

func newCircularList(maxSize int) *circularList {
	return &circularList{nil, nil, 0, maxSize}
}

// find find key from list
func (c *circularList) find(key interface{}) *node {
	prev := c.head
	for i := 0; i < c.size; i++ {
		if prev.key == key {
			return prev
		}
		prev = prev.next
	}
	return nil
}

func (c circularList) hasKey(key interface{}) bool {
	return c.find(key) != nil
}

func (c *circularList) insert(key interface{}, value interface{}) error {
	if c.size == c.capacity {
		return errors.New("capacity is full")
	}

	newNode := &node{key, value, nil, nil}
	if c.size == 0 {
		newNode.next = newNode
		newNode.prev = newNode
		c.head = newNode
		c.tail = newNode
		c.size++
		return nil
	}

	node := c.find(key)
	if node != nil {
		node.value = value
		return nil
	}

	newNode.next = c.head
	newNode.prev = c.tail

	c.tail.next = newNode
	if c.head == c.tail {
		c.head.next = newNode
	}

	c.tail = newNode
	c.head.prev = c.tail

	c.size++

	return nil
}

func (c *circularList) remove(key interface{}) {
	node := c.find(key)
	if node == nil {
		return
	}

	if c.size == 1 {
		c.head = nil
		c.tail = nil
		c.size--
		return
	}

	if node == c.head {
		c.head = c.head.next
	}

	if node == c.tail {
		c.tail = c.tail.prev
	}

	node.next.prev = node.prev
	node.prev.next = node.next

	c.size--
}

func (c *circularList) isFull() bool {
	return c.size == c.capacity
}

func (c *circularList) print() {
	if c.size == 0 {
		fmt.Println(nil)
	}
	ptr := c.head
	for i := 0; i < c.size; i++ {
		fmt.Println(ptr.key, ptr.value, ptr.prev.key, ptr.next.key)
		ptr = ptr.next
	}
}
