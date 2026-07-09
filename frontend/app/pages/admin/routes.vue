<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <h1 class="text-2xl font-bold text-gray-900">Manage Routes</h1>
      <UButton color="primary" icon="i-lucide-plus" @click="openCreateModal">Add Route</UButton>
    </div>

    <UCard :ui="{ body: { padding: '!p-0' } }">
      <UTable :columns="columns" :rows="store.routes">
        <template #actions-data="{ row }">
          <div class="flex gap-2">
            <UButton size="xs" color="gray" variant="ghost" icon="i-lucide-edit" @click="openEditModal(row)" />
            <UButton size="xs" color="red" variant="ghost" icon="i-lucide-trash" @click="deleteRoute(row.id)" />
          </div>
        </template>
      </UTable>
    </UCard>

    <UModal v-model="isModalOpen">
      <UCard :ui="{ ring: '', divide: 'divide-y divide-gray-100' }">
        <template #header>
          <div class="flex items-center justify-between">
            <h3 class="text-base font-semibold leading-6 text-gray-900">
              {{ isEditing ? 'Edit Route' : 'Add New Route' }}
            </h3>
            <UButton color="gray" variant="ghost" icon="i-heroicons-x-mark-20-solid" class="-my-1" @click="isModalOpen = false" />
          </div>
        </template>

        <form @submit.prevent="saveRoute" class="space-y-4">
          <div class="grid grid-cols-2 gap-4">
            <UFormField label="Origin City" required>
              <UInput v-model="form.originCity" required placeholder="e.g. Jakarta" />
            </UFormField>
            <UFormField label="Origin Pool" required>
              <UInput v-model="form.originPool" required placeholder="e.g. Lebak Bulus" />
            </UFormField>
          </div>
          <div class="grid grid-cols-2 gap-4">
            <UFormField label="Destination City" required>
              <UInput v-model="form.destinationCity" required placeholder="e.g. Bandung" />
            </UFormField>
            <UFormField label="Destination Pool" required>
              <UInput v-model="form.destinationPool" required placeholder="e.g. Cihampelas" />
            </UFormField>
          </div>

          <div class="flex justify-end gap-3 mt-6">
            <UButton color="gray" variant="ghost" @click="isModalOpen = false">Cancel</UButton>
            <UButton type="submit" color="primary">Save Route</UButton>
          </div>
        </form>
      </UCard>
    </UModal>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useMockDataStore } from '~/stores/mockData'

definePageMeta({ layout: 'admin' })

const store = useMockDataStore()

const columns = [
  { key: 'id', label: 'ID' },
  { key: 'originCity', label: 'Origin City' },
  { key: 'originPool', label: 'Origin Pool' },
  { key: 'destinationCity', label: 'Destination City' },
  { key: 'destinationPool', label: 'Destination Pool' },
  { key: 'actions', label: 'Actions' }
]

const isModalOpen = ref(false)
const isEditing = ref(false)
const form = ref({
  id: 0,
  originCity: '',
  originPool: '',
  destinationCity: '',
  destinationPool: ''
})

const resetForm = () => {
  form.value = { id: 0, originCity: '', originPool: '', destinationCity: '', destinationPool: '' }
  isEditing.value = false
}

const openCreateModal = () => {
  resetForm()
  isModalOpen.value = true
}

const openEditModal = (route: any) => {
  form.value = { ...route }
  isEditing.value = true
  isModalOpen.value = true
}

const saveRoute = () => {
  if (isEditing.value) {
    store.updateRoute(form.value.id, form.value)
  } else {
    store.addRoute(form.value)
  }
  isModalOpen.value = false
}

const deleteRoute = (id: number) => {
  if (confirm('Are you sure you want to delete this route? Associated schedules will also be deleted.')) {
    store.deleteRoute(id)
  }
}
</script>
