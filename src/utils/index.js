export function deDupe(arr) {
  const unique = []
  const flags = {}
  for (const e of arr) {
    if (!flags[e.name]) {
      flags[e.name] = true
      unique.push(e)
    }
  }
  return unique
}

export function UpperFirst(s) {
  if (typeof s !== 'string') return ''
  return s.charAt(0).toUpperCase() + s.slice(1)
}
