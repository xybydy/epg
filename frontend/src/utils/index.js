export function deDupe(array) {
  const unique = []
  const flags = {}
  array.forEach((value) => {
    if (!flags[value.name]) {
      flags[value.name] = true
      unique.push(value)
    }
  })
  return unique
}
