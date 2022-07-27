<script setup>
import Header from "@/components/Header.vue";
import RoutesWithTransition from "@/router/RoutesWithTransition.vue";
</script>

<template>
  <main>
    <!-- Navigation bar -->
    <Header />
    <!-- Router view -->
    <RoutesWithTransition />
  </main>
</template>

<style>
@import "./assets/base.css";
</style>
