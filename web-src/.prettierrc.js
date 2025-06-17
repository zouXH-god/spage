module.exports = {
  semi: true,
  singleQuote: false,
  trailingComma: "all",
  printWidth: 100,
  tabWidth: 2,
  jsxSingleQuote: false,
  bracketSpacing: true,
  arrowParens: "always",
  importOrder: ["^react", "^[a-z]", "^@", "^[./]"],
  importOrderSeparation: true,
  importOrderSortSpecifiers: true,
  plugins: [
    "prettier-plugin-tailwindcss",
    "@trivago/prettier-plugin-sort-imports"
  ],
};