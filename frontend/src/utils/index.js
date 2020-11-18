export function deDupe(arr) {
  const unique = []
  const flags = {}
  arr.forEach((e) => {
    if (!flags[e.name]) {
      flags[e.name] = true
      unique.push(e)
    }
  })
  return unique
}
