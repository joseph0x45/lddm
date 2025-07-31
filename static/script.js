document.addEventListener('alpine:init', () => {
  Alpine.store('modals_state', Alpine.reactive({
    show_create_product_modal: false,
    show_view_cart_modal: false,
  }))
  Alpine.store('cart', Alpine.reactive({}))
});

function product_card() {
  return {
    quantity: 0,
    increase_quantity() {
      this.quantity++
      dataset = this.$root.dataset
      current_product = {
        id: dataset.id,
        name: dataset.name,
        variant: dataset.variant,
        price: dataset.price,
        quantity: this.quantity
      }
      this.$store.cart[current_product.id] = current_product
    },
    decrease_quantity() {
      dataset = this.$root.dataset
      if (this.quantity == 0) {
        return
      }
      this.quantity--
      if (this.quantity == 0) {
        delete this.$store.cart[dataset.id]
        return
      }
      current_product = {
        id: dataset.id,
        name: dataset.name,
        variant: dataset.variant,
        price: dataset.price,
        quantity: this.quantity
      }
      this.$store.cart[current_product.id] = current_product
    }
  }
}

function create_product_modal() {
  return {
    name: "",
    variant: "mini",
    price: 0,
    description: "",
    image: "",
    in_stock: 0,
    toggle_modal() {
      this.$store.modals_state.show_create_product_modal = !this.$store.modals_state.show_create_product_modal
      window.location.reload()
    },
    async create_product() {
      try {
        if (this.name == "") {
          alert("Name can not be empty")
          return
        }
        if (this.image == "") {
          this.image = "https://picsum.photos/200/300"
        }
        const response = await fetch(
          "/api/products", {
          method: "POST",
          body: JSON.stringify({
            name: this.name,
            variant: this.variant,
            price: parseInt(this.price),
            description: this.description,
            image: this.image,
            in_stock: parseInt(this.in_stock)
          })
        })
        if (response.status != 201) {
          const error_message = await response.text()
          throw Error(error_message)
        }
        alert('Product created successfully')
        this.name = ""
        this.variant = "mini"
        this.price = "0"
        this.description = ""
        this.image = ""
        this.in_stock = "0"
      } catch (error) {
        alert(`Error while creating product: ${error}`)
      }
    }
  }
}

function layout_handler() {
  return {
    open_sidebar: false,
    toggle_sidebar() {
      this.open_sidebar = !this.open_sidebar
    },
    toggle_create_product_modal() {
      this.toggle_sidebar()
      this.$store.modals_state.show_create_product_modal = !this.$store.modals_state.show_create_product_modal
    },
    toggle_view_cart_modal() {
      this.toggle_sidebar()
      this.$store.modals_state.show_view_cart_modal = !this.$store.modals_state.show_view_cart_modal
    },
  }
}

function view_cart_modal() {
  return {
    customer_name: "",
    customer_phone: "",
    customer_address: "",
    discount: "0",
    get_cart_data() {
      return Object.values(this.$store.cart)
    },
    toggle_view_cart_modal() {
      this.$store.modals_state.show_view_cart_modal = !this.$store.modals_state.show_view_cart_modal
    },
    get total() {
      let total = 0
      this.get_cart_data().forEach(element => {
        total += element.quantity * element.price
      });
      return total
    },
    get subtotal() {
      if (this.discount >= this.total){
        this.discount = this.total
        return 0
      }
      return this.total - this.discount
    }

  }
}
