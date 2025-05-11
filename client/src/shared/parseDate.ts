export default function parseDate(inputDate?: string): string {
  if (!inputDate) return ''
  return `${new Date(inputDate).toLocaleDateString()} ${new Date(inputDate).toLocaleTimeString()}`
}
