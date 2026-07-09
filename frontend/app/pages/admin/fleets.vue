<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-bold text-gray-900">Manage Fleets</h1>
      <UButton color="primary" icon="i-lucide-plus" @click="openCreateModal">Add Fleet</UButton>
    </div>

    <UCard :ui="{ body: { padding: '!p-0' } }">
      <UTable :columns="columns" :rows="store.fleets">
        <template #type-data="{ row }">
          <span 
            class="inline-flex items-center px-2 py-0.5 rounded text-xs font-medium"
            :class="row.type === 'Premium' ? 'bg-amber-100 text-amber-800' : 'bg-gray-100 text-gray-800'"
          >
            {{ row.type }}
          </span>
        </template>
        <template #actions-data="{ row }">
          <div class="flex gap-2">
            <UButton size="xs" color="gray" variant="ghost" icon="i-lucide-edit" @click="openEditModal(row)" />
            <UButton size="xs" color="red" variant="ghost" icon="i-lucide-trash" @click="deleteFleet(row.id)" />
          </div>
        </template>
      </UTable>
    </UCard>

    <UModal v-model="isModalOpen">
      <UCard :ui="{ ring: '', divide: 'divide-y divide-gray-100' }">
        <template #header>
          <div class="flex items-center justify-between">
            <h3 class="text-base font-semibold leading-6 text-gray-900">
              {{ isEditing ? 'Edit Fleet' : 'Add New Fleet' }}
            </h3>
            <UButton color="gray" variant="ghost" icon="i-heroicons-x-mark-20-solid" class="-my-1" @click="isModalOpen = false" />
          </div>
        </template>

        <form @submit.prevent="saveFleet" class="space-y-4">
          <UFormField label="Fleet Name" required>
            <UInput v-model="form.name" required placeholder="e.g. MesenShuttle 08" />
          </UFormField>
          
          <div class="grid grid-cols-2 gap-4">
            <UFormField label="Fleet Type" required>
              <USelect v-model="form.type" :options="['Reguler', 'Premium']" required />
            </UFormField>
            <UFormField label="Capacity (Seats)" required>
              <UInput v-model.number="form.capacity" type="number" required min="1" max="50" />
            </UFormField>
          </div>

          <div class="flex justify-end gap-3 mt-6">
            <UButton color="gray" variant="ghost" @click="isModalOpen = false">Cancel</UButton>
            <UButton type="submit" color="primary">Save Fleet</UButton>
          </div>
        </form>
      </UCard>
    </UModal>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useMockDataStore } from '~/stores/mockData'

definePageMeta({ layout: 'admin' })

const store = useMockDataStore()

const columns = [
  { key: 'id', label: 'ID' },
  { key: 'name', label: 'Fleet Name' },
  { key: 'type', label: 'Type' },
  { key: 'capacity', label: 'Capacity' },
  { key: 'actions', label: 'Actions' }
]

const isModalOpen = ref(false)
const isEditing = ref(false)
const form = ref({
  id: 0,
  name: '',
  vehicle: 'Shuttle',
  type: 'Reguler',
  capacity: 12
})

// Auto update capacity based on type for convenience
watch(() => form.value.type, (newType) => {
  if (!isEditing.value) {
    form.value.capacity = newType === 'Premium' ? 10 : 12
  }
})

const resetForm = () => {
  form.value = { id: 0, name: '', vehicle: 'Shuttle', type: 'Reguler', capacity: 12 }
  isEditing.value = false
}

const openCreateModal = () => {
  resetForm()
  isModalOpen.value = true
}

const openEditModal = (fleet: any) => {
  form.value = { ...fleet }
  isEditing.value = true
  isModalOpen.value = true
}

const saveFleet = () => {
  if (isEditing.value) {
    store.updateFleet(form.value.id, form.value)
  } else {
    store.addFleet(form.value)
  }
  isModalOpen.value = false
}

const deleteFleet = (id: number) => {
  if (confirm('Are you sure you want to delete this fleet? Associated schedules will also be deleted.')) {
    store.deleteFleet(id)
  }
}
</script>
