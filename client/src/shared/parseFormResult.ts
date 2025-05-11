import type { FormFieldState } from '@primevue/forms'

export function parseFormResult<T>(states: Record<string, FormFieldState>): T {
  return Object.entries(states).reduce(
    (prev, [key, value]) => ({ ...prev, [key]: value.value }),
    {} as T,
  )
}
