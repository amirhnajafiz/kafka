<template>
  <header class="header">
    <Logo class="logo" />
    <nav class="nav">
      <RouterLink class="select" to="/">Amirhossein Najafizade</RouterLink>
      <div>
        <RouterLink class="link select" to="/">Projects</RouterLink>
        <RouterLink class="link select" to="/">Educations</RouterLink>
        <RouterLink class="link select" to="/">About</RouterLink>
        <RouterLink class="link select" to="/">Contact</RouterLink>
      </div>
    </nav>
  </header>
</template>

<script>
import Logo from './Logo.vue';
export default {
    name: "Header",
    components: { Logo }
};
</script>

<style scoped>
a {
  text-decoration: none;
}

.header {
  width: 100%;
}

.logo {
  width: 40%;
  margin: 100px auto 50px;
}

.nav {
  width: 60%;
  margin: 0 auto;
  padding: 0 0 10px;

  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: space-between;

  border-bottom: 1px solid #ffffff;
}

.link {
  margin: 0 10px;
}

.select:hover {
  color: #aaaaaa;
}
</style>
