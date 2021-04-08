import XCTest
@testable import swiftLearn

final class swiftLearnTests: XCTestCase {
    func testExample() {
        // This is an example of a functional test case.
        // Use XCTAssert and related functions to verify your tests produce the correct
        // results.
        XCTAssertEqual(swiftLearn().text, "Hello, World!")
    }

    static var allTests = [
        ("testExample", testExample),
    ]
}
