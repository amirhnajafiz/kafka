<script setup>
import Header from "@/components/Header.vue";
import RoutesWithTransition from "@/router/RoutesWithTransition.vue";
</script>

<template>
  <!-- Navigation bar -->
  <Header />
  <!-- Router view -->
  <RoutesWithTransition />
</template>

<style>
@import './assets/base.css';
</style>
