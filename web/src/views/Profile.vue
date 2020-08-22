<template>
  <section>
    <b-tabs position="is-centered" class="block">
      <b-tab-item label="Dashboards"></b-tab-item>
      <b-tab-item label="Collections"></b-tab-item>
      <b-tab-item label="Preferences"></b-tab-item>
    </b-tabs>
    <div class="columns">
      <div class="column is-one-quarter has-background">
        <div class="profile-card">
          <b-image
            src="https://avatars3.githubusercontent.com/u/210818?s=460&u=6a66ae45f34c5d7e3db4508b463bc74021087032&v=4"
            rounded
            responsive
            ratio="1by1"
          />
          <div class="mt-5 px-3">
            <h5 class="is-size-5">{{ currentUser.name }}</h5>
            <h5 class="is-size-5 has-text-grey">{{ currentUser.email }}</h5>
            <b-button expanded class="mt-4">
              Edit Profile
            </b-button>
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
  }
  .block {
  }
}
</style>

<script>
import { mapState } from "vuex"
export default {
  name: "ProfileView",
  computed: {
    ...mapState("authentication", ["currentUser"])
  }
}
</script>
