module.exports = {
  root: true,
  env: {
    node: true,
    es6: true,
  },
  extends: [
    'eslint:recommended',
    'plugin:vue/base',
    'plugin:vue/vue3-essential',
    'plugin:vue/vue3-recommended',
    'plugin:vue/vue3-strongly-recommended',
    'prettier',
    'prettier/vue',
    'prettier/unicorn',
    'prettier/standard',
    'plugin:prettier/recommended',
    '@vue/prettier',
    'plugin:unicorn/recommended',
  ],
  parserOptions: {
    parser: 'babel-eslint',
  },
  rules: {
    'no-console': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-debugger': process.env.NODE_ENV === 'production' ? 'warn' : 'off',
    'no-unused-vars': 1,
    'unicorn/filename-case': [
      'warn',
      {
        case: 'pascalCase',
        ignore: ['.*.js', 'babel.config.js'],
      },
    ],
    'unicorn/consistent-function-scoping': 'off',
    // 'vue/no-v-model-argument': 1
  },
}