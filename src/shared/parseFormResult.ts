import type { FormFieldState } from '@primevue/forms'

export function parseFormResult<T>(states: Record<string, FormFieldState>): Record<string, T> {
  return Object.entries(states).reduce(
    (prev, [key, value]) => ({ ...prev, [key]: value.value }),
    {},
  )
}
