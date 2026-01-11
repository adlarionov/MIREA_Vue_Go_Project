interface Option {
  label: string
  value: string
}

export const getOptionsFromData = <T extends Record<keyof T, T[keyof T]>>(
  data: T[],
  keyForValue: keyof T,
  keyForLabel: keyof T,
): Option[] => {
  return data.reduce<Option[]>((dataItem, currentItem) => {
    dataItem.push({ label: currentItem[keyForLabel], value: currentItem[keyForValue] })

    return dataItem
  }, [])
}
