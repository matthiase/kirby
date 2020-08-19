<template>
  <div class="container">
    <Navbar />
    <router-view />
  </div>
</template>

<script>
import Navbar from "@/components/Navbar"
export default {
  name: "ApplicationLayout",
  components: {
    Navbar
  }
}
</script>
