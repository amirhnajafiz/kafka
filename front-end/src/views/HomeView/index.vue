<script setup>
import { useInformationStore } from "@/stores/information";
</script>

<template>
  <main>
    <h1>
      {{ useInformationStore().getFirstName + " " + useInformationStore().getLastName }}
    </h1>
  </main>
</template>
