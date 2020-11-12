export function deDupe(arr) {
  return arr.filter((v, i, a) => a.findIndex(t => t.chan_name === v.chan_name) === i)
}
