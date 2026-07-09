<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-bold text-gray-900">Manage Schedules</h1>
      <UButton color="primary" icon="i-lucide-plus" @click="openCreateModal">Add Schedule</UButton>
    </div>

    <UCard :ui="{ body: { padding: '!p-0' } }">
      <UTable :columns="columns" :rows="tableRows">
        <template #departureTime-data="{ row }">
          {{ new Date(row.departureTime).toLocaleString('id-ID', { dateStyle: 'medium', timeStyle: 'short' }) }}
        </template>
        <template #price-data="{ row }">
          Rp {{ row.price.toLocaleString() }}
        </template>
        <template #actions-data="{ row }">
          <div class="flex gap-2">
            <UButton size="xs" color="gray" variant="ghost" icon="i-lucide-edit" @click="openEditModal(row)" />
            <UButton size="xs" color="red" variant="ghost" icon="i-lucide-trash" @click="deleteSchedule(row.id)" />
          </div>
        </template>
      </UTable>
    </UCard>

    <UModal v-model="isModalOpen">
      <UCard :ui="{ ring: '', divide: 'divide-y divide-gray-100' }">
        <template #header>
          <div class="flex items-center justify-between">
            <h3 class="text-base font-semibold leading-6 text-gray-900">
              {{ isEditing ? 'Edit Schedule' : 'Add New Schedule' }}
            </h3>
            <UButton color="gray" variant="ghost" icon="i-heroicons-x-mark-20-solid" class="-my-1" @click="isModalOpen = false" />
          </div>
        </template>

        <form @submit.prevent="saveSchedule" class="space-y-4">
          <UFormField label="Route" required>
            <USelect 
              v-model.number="form.routeId" 
              :options="store.routes.map(r => ({ label: `${r.originCity} (${r.originPool}) → ${r.destinationCity} (${r.destinationPool})`, value: r.id }))" 
              required 
            />
          </UFormField>
          
          <UFormField label="Fleet" required>
            <USelect 
              v-model.number="form.fleetId" 
              :options="store.fleets.map(f => ({ label: `${f.name} - ${f.type} (${f.capacity} seats)`, value: f.id }))" 
              required 
            />
          </UFormField>
          
          <div class="grid grid-cols-2 gap-4">
            <UFormField label="Departure Time" required>
              <UInput v-model="form.departureTime" type="datetime-local" required />
            </UFormField>
            
            <UFormField label="Price (Rp)" required>
              <UInput v-model.number="form.price" type="number" required min="0" />
            </UFormField>
          </div>

          <div class="flex justify-end gap-3 mt-6">
            <UButton color="gray" variant="ghost" @click="isModalOpen = false">Cancel</UButton>
            <UButton type="submit" color="primary">Save Schedule</UButton>
          </div>
        </form>
      </UCard>
    </UModal>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useMockDataStore } from '~/stores/mockData'

definePageMeta({ layout: 'admin' })

const store = useMockDataStore()

const columns = [
  { key: 'id', label: 'ID' },
  { key: 'route', label: 'Route' },
  { key: 'fleet', label: 'Fleet' },
  { key: 'departureTime', label: 'Departure Time' },
  { key: 'price', label: 'Price' },
  { key: 'actions', label: 'Actions' }
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
  // Format for datetime-local
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
  // Convert local datetime to ISO string for store
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
