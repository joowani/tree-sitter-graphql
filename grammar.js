/**
 * @file GraphQL grammar for tree-sitter
 * @author Joohwan Oh <joohwan.oh@outlook.com>
 * @license MIT
 */

/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "graphql",

  rules: {
    // TODO: add the actual grammar rules
    source_file: $ => "hello"
  }
});
