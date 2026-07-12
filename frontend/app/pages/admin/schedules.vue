<template>
  <div class="space-y-6 pb-12">
    <!-- Header -->
    <AdminPageHeader 
      title="Manage Schedules" 
      description="Create and manage departure schedules and pricing." 
      button-label="Add Schedule" 
      button-icon="i-lucide-calendar-plus" 
      @action="openCreateModal" 
    />

    <UCard class="ring-1 ring-gray-100 rounded-2xl shadow-sm" :ui="{ body: '!p-0' }">
      <UTable :columns="columns" :rows="tableRows" :ui="{ td: 'whitespace-nowrap transition-colors hover:bg-gray-50/50' }">
        <template #departureTime-cell="{ row }">
          <div class="flex flex-col">
            <span class="font-medium text-gray-900">{{ new Date(row.original.departureTime as string).toLocaleDateString('id-ID', { dateStyle: 'medium' }) }}</span>
            <span class="text-xs text-gray-500">{{ new Date(row.original.departureTime as string).toLocaleTimeString('id-ID', { timeStyle: 'short' }) }}</span>
          </div>
        </template>
        <template #price-cell="{ row }">
          <span class="font-semibold text-gray-900">Rp {{ ((row.original.price as number) || 0).toLocaleString() }}</span>
        </template>
        <template #actions-cell="{ row }">
          <div class="flex gap-2">
            <UButton size="xs" color="neutral" variant="ghost" icon="i-lucide-edit" @click="openEditModal(row.original)" />
            <UButton size="xs" color="error" variant="ghost" icon="i-lucide-trash" @click="deleteSchedule(row.original.id as number)" />
          </div>
        </template>
        <template #empty-state>
          <div class="flex flex-col items-center justify-center py-12 px-4 text-center">
            <UIcon name="i-lucide-calendar-x" class="w-12 h-12 text-gray-300 mb-4" />
            <p class="text-sm font-medium text-gray-900">No schedules found</p>
            <p class="text-sm text-gray-500 mt-1">Get started by creating a new departure schedule.</p>
          </div>
        </template>
      </UTable>
    </UCard>

    <AdminScheduleModal v-model:open="isModalOpen" v-model:form="form" :is-editing="isEditing" @save="saveSchedule" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useMockDataStore } from '~/stores/mockData'

definePageMeta({ layout: 'admin' })

const store = useMockDataStore()

const columns = [
  { accessorKey: 'id', header: 'ID' },
  { accessorKey: 'route', header: 'Route' },
  { accessorKey: 'fleet', header: 'Fleet' },
  { accessorKey: 'departureTime', header: 'Departure Time' },
  { accessorKey: 'price', header: 'Price' },
  { accessorKey: 'actions', header: 'Actions' }
]

const tableRows = computed(() => {
  return store.schedules.map(schedule => {
    const details = store.getScheduleDetails(schedule.id)
    return {
      ...schedule,
      route: details?.route ? `${details.route.originCity} → ${details.route.destinationCity}` : 'Unknown Route',
      fleet: details?.fleet ? `${details.fleet.name} (${details.fleet.type})` : 'Unknown Fleet'
    }
  }).sort((a, b) => new Date(b.departureTime).getTime() - new Date(a.departureTime).getTime())
})

const isModalOpen = ref(false)
const isEditing = ref(false)
const form = ref({
  id: 0,
  routeId: store.routes[0]?.id || 1,
  fleetId: store.fleets[0]?.id || 1,
  departureTime: new Date().toISOString().slice(0, 16),
  price: 150000
})

const resetForm = () => {
  form.value = { 
    id: 0, 
    routeId: store.routes[0]?.id || 1, 
    fleetId: store.fleets[0]?.id || 1, 
    departureTime: new Date().toISOString().slice(0, 16), 
    price: 150000 
  }
  isEditing.value = false
}

const openCreateModal = () => {
  resetForm()
  isModalOpen.value = true
}

const openEditModal = (schedule: any) => {
  let formattedTime = ''
  if (schedule.departureTime) {
    const dateObj = new Date(schedule.departureTime)
    formattedTime = new Date(dateObj.getTime() - (dateObj.getTimezoneOffset() * 60000)).toISOString().slice(0, 16)
  }
  
  form.value = { 
    ...schedule,
    departureTime: formattedTime
  }
  isEditing.value = true
  isModalOpen.value = true
}

const saveSchedule = () => {
  const isoTime = new Date(form.value.departureTime).toISOString()
  
  if (isEditing.value) {
    store.updateSchedule(form.value.id, { ...form.value, departureTime: isoTime })
  } else {
    store.addSchedule({ ...form.value, departureTime: isoTime })
  }
  isModalOpen.value = false
}

const deleteSchedule = (id: number) => {
  if (confirm('Are you sure you want to delete this schedule?')) {
    store.deleteSchedule(id)
  }
}
</script>
