<template>
  <div class="space-y-6 pb-12">
    <!-- Header -->
    <AdminPageHeader 
      title="Manage Fleets" 
      description="Monitor and manage all operational vehicles." 
      button-label="Add Fleet" 
      button-icon="i-lucide-plus" 
      @action="openCreateModal" 
    />

    <!-- Table Card -->
    <UCard class="ring-1 ring-gray-100 rounded-2xl shadow-sm" :ui="{ body: '!p-0' }">
      <UTable :columns="columns" :rows="store.fleets" :ui="{ td: 'whitespace-nowrap transition-colors hover:bg-gray-50/50' }">
        <template #name-cell="{ row }">
          <div class="flex items-center gap-3">
            <div class="w-8 h-8 rounded-full bg-gray-100 flex items-center justify-center text-gray-500">
              <UIcon name="i-lucide-bus" class="w-4 h-4" />
            </div>
            <span class="font-medium text-gray-900">{{ row.original.name }}</span>
          </div>
        </template>
        <template #type-cell="{ row }">
          <UBadge :color="row.original.type === 'Premium' ? 'warning' : 'neutral'" variant="subtle" size="sm">
            {{ row.original.type }}
          </UBadge>
        </template>
        <template #capacity-cell="{ row }">
          <span class="text-gray-600">{{ row.original.capacity }} Seats</span>
        </template>
        <template #platNumber-cell="{ row }">
          <span class="font-mono text-sm text-gray-700 bg-gray-100/80 border border-gray-200 px-2.5 py-1 rounded-md">{{ row.original.platNumber || '-' }}</span>
        </template>
        <template #actions-cell="{ row }">
          <div class="flex gap-2">
            <UButton size="xs" color="neutral" variant="ghost" icon="i-lucide-edit" @click="openEditModal(row.original)" />
            <UButton size="xs" color="error" variant="ghost" icon="i-lucide-trash" @click="deleteFleet(row.original.id as number)" />
          </div>
        </template>
        <template #empty-state>
          <div class="flex flex-col items-center justify-center py-12 px-4 text-center">
            <UIcon name="i-lucide-bus" class="w-12 h-12 text-gray-300 mb-4" />
            <p class="text-sm font-medium text-gray-900">No fleets found</p>
            <p class="text-sm text-gray-500 mt-1">Get started by adding a new vehicle.</p>
          </div>
        </template>
      </UTable>
    </UCard>

    <AdminFleetModal v-model:open="isModalOpen" v-model:form="form" :is-editing="isEditing" @save="saveFleet" />
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { useMockDataStore } from '~/stores/mockData'

definePageMeta({ layout: 'admin' })

const store = useMockDataStore()

const columns = [
  { accessorKey: 'id', header: 'ID' },
  { accessorKey: 'name', header: 'Fleet Name' },
  { accessorKey: 'platNumber', header: 'Plat Number' },
  { accessorKey: 'type', header: 'Type' },
  { accessorKey: 'capacity', header: 'Capacity' },
  { accessorKey: 'actions', header: 'Actions' }
]

const isModalOpen = ref(false)
const isEditing = ref(false)
const form = ref({
  id: 0,
  name: '',
  platNumber: '',
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
  form.value = { id: 0, name: '', platNumber: '', vehicle: 'Shuttle', type: 'Reguler', capacity: 12 }
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
