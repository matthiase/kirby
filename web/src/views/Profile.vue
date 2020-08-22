<template>
  <section>
    <b-tabs position="is-centered" class="block">
      <b-tab-item label="Dashboards"></b-tab-item>
      <b-tab-item label="Collections"></b-tab-item>
      <b-tab-item label="Preferences"></b-tab-item>
    </b-tabs>
    <div class="columns">
      <div class="column is-one-quarter">
        <div class="profile-card">
          <b-image
            class="avatar-image"
            src="https://avatars3.githubusercontent.com/u/210818?s=460&u=6a66ae45f34c5d7e3db4508b463bc74021087032&v=4"
            responsive
            is-1by1
            rounded
          />
          <div class="mt-5 px-3 has-text-centered">
            <form v-if="isEditing" @submit.prevent="handleSaveProfileClick" class="px-2">
              <b-field>
                <b-input v-model="user.name" placeholder="Full name" />
              </b-field>
              <b-field>
                <b-input type="email" v-model="user.email" placeholder="Email address" />
              </b-field>
              <div class="columns mt-4">
                <div class="column">
                  <b-button type="is-primary" expanded>
                    Save
                  </b-button>
                </div>
                <div class="column" @click="handleCancelProfileClick">
                  <b-button expanded>
                    Cancel
                  </b-button>
                </div>
              </div>
            </form>
            <div v-else>
              <h5 class="is-size-5">{{ currentUser.name }}</h5>
              <h6 class="is-size-6 has-text-grey">{{ currentUser.email }}</h6>
              <b-button expanded class="mt-5" @click="handleEditProfileClick">
                Edit Profile
              </b-button>
            </div>
          </div>
        </div>
      </div>
      <div class="column">
        &nbsp;
      </div>
    </div>
  </section>
</template>

<style lang="scss" scoped>
.profile-card {
  @media (min-width: 768px) {
    margin-top: -90px;
    position: relative;
    z-index: 100;
  }
  .avatar-image {
    margin: 0 20%;
  }
}
</style>

<script>
import { mapState } from "vuex"
export default {
  name: "ProfileView",
  data() {
    return {
      isEditing: false,
      user: {
        name: "",
        email: ""
      }
    }
  },
  computed: {
    ...mapState("authentication", ["currentUser"])
  },
  methods: {
    handleEditProfileClick() {
      this.user.name = this.currentUser.name
      this.user.email = this.currentUser.email
      this.isEditing = true
    },
    handleCancelProfileClick() {
      this.isEditing = false
    }
  }
}
</script>
