<template>
  <div class="container">
    <b-navbar>
      <template slot="brand">
        <b-navbar-item tag="router-link" :to="{ path: '/' }">
          <img src="../../assets/logo.svg" alt="Kirby Logo" class="image is-32x32" />
        </b-navbar-item>
      </template>
      <template slot="end">
        <b-navbar-item tag="div">
          <div class="buttons">
            <b-button type="is-info" tag="router-link" to="/login">
              Sign in
            </b-button>
          </div>
        </b-navbar-item>
      </template>
    </b-navbar>
    <router-view />
  </div>
</template>

<script>
export default {
  name: 'PublicLayout'
}
</script>
