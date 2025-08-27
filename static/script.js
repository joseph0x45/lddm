function app() {
  return {
    search: '',
    groups: [],
    products: [],
    selectedGroup: null,
    async init() {
      try {
        const groups = await fetch("/api/groups")
        this.groups = await groups.json()

        const products = await fetch("/api/products")
        this.products = await products.json()

        if (this.groups.length) {
          this.selectedGroup = this.groups[0].id
        }
      } catch (err) {
        alert("Error fetching data")
        console.error(`Error fetching data ${err}`)
      }
    },
    selectGroup(id){
      this.selectedGroup = id
    },
    get filteredGroups() {
      if (this.search == "") {
        return this.groups
      }
      return this.groups.filter(group => {
        return group.name.toLowerCase().includes(this.search.toLowerCase());
      })
    }
  }
}
