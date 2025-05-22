import XCTest
import SwiftTreeSitter
import TreeSitterVba

final class TreeSitterVbaTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_vba())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading no grammar")
    }
}
