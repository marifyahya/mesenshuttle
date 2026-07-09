<template>
  <div>
    <h1 class="text-2xl font-bold mb-6 text-gray-900">Dashboard Overview</h1>
    
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
      <UCard :ui="{ body: { padding: 'p-6' } }">
        <div class="flex items-center gap-4">
          <div class="p-3 bg-blue-50 rounded-xl text-blue-600 border border-blue-100">
            <UIcon name="i-lucide-map" class="w-8 h-8" />
          </div>
          <div>
            <p class="text-sm font-medium text-gray-500">Total Routes</p>
            <p class="text-3xl font-bold text-gray-900">{{ store.routes.length }}</p>
          </div>
        </div>
      </UCard>
      
      <UCard :ui="{ body: { padding: 'p-6' } }">
        <div class="flex items-center gap-4">
          <div class="p-3 bg-purple-50 rounded-xl text-purple-600 border border-purple-100">
            <UIcon name="i-lucide-bus" class="w-8 h-8" />
          </div>
          <div>
            <p class="text-sm font-medium text-gray-500">Total Fleets</p>
            <p class="text-3xl font-bold text-gray-900">{{ store.fleets.length }}</p>
          </div>
        </div>
      </UCard>
      
      <UCard :ui="{ body: { padding: 'p-6' } }">
        <div class="flex items-center gap-4">
          <div class="p-3 bg-green-50 rounded-xl text-green-600 border border-green-100">
            <UIcon name="i-lucide-calendar-clock" class="w-8 h-8" />
          </div>
          <div>
            <p class="text-sm font-medium text-gray-500">Active Schedules</p>
            <p class="text-3xl font-bold text-gray-900">{{ store.schedules.length }}</p>
          </div>
        </div>
      </UCard>
    </div>

    <h2 class="text-xl font-bold mb-4 text-gray-900">Quick Actions</h2>
    <div class="flex flex-wrap gap-4">
      <UButton to="/admin/routes" color="gray" variant="solid" icon="i-lucide-plus" size="lg">Add New Route</UButton>
      <UButton to="/admin/fleets" color="gray" variant="solid" icon="i-lucide-plus" size="lg">Add New Fleet</UButton>
      <UButton to="/admin/schedules" color="primary" variant="solid" icon="i-lucide-plus" size="lg">Schedule a Trip</UButton>
    </div>
  </div>
</template>

<script setup lang="ts">
import { useMockDataStore } from '~/stores/mockData'

definePageMeta({
  layout: 'admin'
})

const store = useMockDataStore()
</script>
