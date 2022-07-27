<script setup>
import Header from "@/components/Header.vue";
import RoutesWithTransition from "@/router/RoutesWithTransition.vue";
import Footer from "./components/Footer.vue";
</script>

<template>
  <main>
    <!-- Navigation bar -->
    <Header />
    <!-- Router view -->
    <RoutesWithTransition />
    <!-- Footer -->
    <Footer />
  </main>
</template>

<style>
@import "./assets/base.css";
</style>
