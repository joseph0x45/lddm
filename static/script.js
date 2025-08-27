document.addEventListener('alpine:init', () => {
  Alpine.data('app', () => ({
    data: {},
    groups: [],
    products: [],
    search: '',
    selectedGroup: null,
    async init() {
      const data = await fetch("/api/data").then(async (res) => await res.json())
      this.data = data
      this.groups = data.groups
      this.selectedGroup = data.groups[0].id
      this.products = data.products
    },
    get filteredGroups() {
      if (this.search == '') {
        return this.groups
      }
      return this.groups.filter(group => group.name.toLowerCase().includes(this.search.toLowerCase()));
    },
    selectGroup(id) {
      this.selectedGroup = id
    },
    get filteredProducts() {
      return this.products.filter(p =>
        p.group === this.selectedGroup
      )
    }
  }))
})
