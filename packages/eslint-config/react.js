const { resolve } = require("node:path");

const project = resolve(process.cwd(), "tsconfig.json");

module.exports = {
  extends: [
    "@vercel/style-guide/eslint/browser",
    "@vercel/style-guide/eslint/typescript",
    "@vercel/style-guide/eslint/react",
  ].map(require.resolve),
  parserOptions: {
    project,
  },
  plugins: ["only-warn", "simple-import-sort"],
  globals: {
    JSX: true,
  },
  settings: {
    "import/resolver": {
      typescript: {
        project,
      },
    },
  },
  ignorePatterns: ["node_modules/", "dist/", ".eslintrc.js", "**/*.css"],
  overrides: [
    {
      files: ["*.config.js"],
      env: {
        node: true,
      },
    },
  ],
  rules: {
    "import/order": "off",
    "import/no-default-export": "off",
    "simple-import-sort/imports": "error",
    "simple-import-sort/exports": "error",
  },
};
