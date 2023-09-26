export const shortcuts = [
  {
    text: "最近一周",
    value: () => {
      const end = new Date()
      const start = new Date()
      // 将毫秒部分设置为零
      start.setMilliseconds(0);
      end.setMilliseconds(0);
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 7)
      return [start, end]
    }
  },
  {
    text: "最近一月",
    value: () => {
      const end = new Date()
      const start = new Date()
      // 将毫秒部分设置为零
      start.setMilliseconds(0);
      end.setMilliseconds(0);
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 30)
      return [start, end]
    }
  },
  {
    text: "最近三月",
    value: () => {
      const end = new Date()
      const start = new Date()
      // 将毫秒部分设置为零
      start.setMilliseconds(0);
      end.setMilliseconds(0);
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 90)
      return [start, end]
    }
  },
  {
    text: "最近一年",
    value: () => {
      const end = new Date()
      const start = new Date()
      // 将毫秒部分设置为零
      start.setMilliseconds(0);
      end.setMilliseconds(0);
      start.setTime(start.getTime() - 3600 * 1000 * 24 * 365)
      return [start, end]
    }
  }
]


export const expireDateshortcuts = [
  {
    text: "一月",
    value: () => {
      const d = new Date()
      d.setMilliseconds(0)
      d.setTime(d.getTime() + 3600 * 1000 * 24 * 31)
      return d
    }
  },
  {
    text: "一年",
    value: () => {
      const d = new Date()
      d.setMilliseconds(0)
      d.setTime(d.getTime() + 3600 * 1000 * 24 * 365)
      return d
    }
  },
  {
    text: "永久",
    value: () => {
      const d = new Date()
      d.setMilliseconds(0)
      d.setTime(d.getTime() + 3600 * 1000 * 24 * 36500)
      return d
    }
  },
]
