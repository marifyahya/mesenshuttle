<template>
  <div class="space-y-6 pb-12">
    <!-- Header -->
    <AdminPageHeader 
      title="Manage Routes" 
      description="Define and manage operational routes across cities." 
      button-label="Add Route" 
      button-icon="i-lucide-plus" 
      @action="openCreateModal" 
    />

    <!-- Table Card -->
    <UCard class="ring-1 ring-gray-100 rounded-2xl shadow-sm" :ui="{ body: '!p-0' }">
      <UTable :columns="columns" :rows="store.routes" :ui="{ td: 'whitespace-nowrap transition-colors hover:bg-gray-50/50' }">
        <template #originCity-cell="{ row }">
          <div class="flex flex-col">
            <span class="font-medium text-gray-900">{{ row.original.originCity }}</span>
            <span class="text-xs text-gray-500">Pool: {{ row.original.originPool }}</span>
          </div>
        </template>
        <template #destinationCity-cell="{ row }">
          <div class="flex flex-col">
            <span class="font-medium text-gray-900">{{ row.original.destinationCity }}</span>
            <span class="text-xs text-gray-500">Pool: {{ row.original.destinationPool }}</span>
          </div>
        </template>
        <template #actions-cell="{ row }">
          <div class="flex gap-2">
            <UButton size="xs" color="neutral" variant="ghost" icon="i-lucide-edit" @click="openEditModal(row.original)" />
            <UButton size="xs" color="error" variant="ghost" icon="i-lucide-trash" @click="deleteRoute(row.original.id as number)" />
          </div>
        </template>
        <template #empty-state>
          <div class="flex flex-col items-center justify-center py-12 px-4 text-center">
            <UIcon name="i-lucide-map-x" class="w-12 h-12 text-gray-300 mb-4" />
            <p class="text-sm font-medium text-gray-900">No routes found</p>
            <p class="text-sm text-gray-500 mt-1">Get started by creating a new operational route.</p>
          </div>
        </template>
      </UTable>
    </UCard>

    <AdminRouteModal v-model:open="isModalOpen" v-model:form="form" :is-editing="isEditing" @save="saveRoute" />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useMockDataStore } from '~/stores/mockData'

definePageMeta({ layout: 'admin' })

const store = useMockDataStore()

const columns = [
  { accessorKey: 'id', header: 'ID' },
  { accessorKey: 'originCity', header: 'Origin' },
  { accessorKey: 'destinationCity', header: 'Destination' },
  { accessorKey: 'actions', header: 'Actions' }
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
