<template>
  <UModal v-model:open="isOpen">
    <template #content>
      <UCard class="ring-0 divide-y divide-gray-100 rounded-3xl shadow-2xl overflow-hidden bg-white/80 backdrop-blur-xl">
        <template #header>
          <div class="flex items-start justify-between px-2 pt-2">
            <div class="flex items-center gap-4">
              <div class="p-3.5 bg-linear-to-br from-primary-500 to-primary-600 text-white rounded-2xl shadow-lg shadow-primary-500/30 ring-1 ring-primary-600/20">
                <UIcon :name="isEditing ? 'i-lucide-edit-3' : 'i-lucide-map-pin-plus'" class="w-6 h-6" />
              </div>
              <div>
                <h3 class="text-xl font-bold text-gray-900 tracking-tight">
                  {{ isEditing ? 'Edit Route Details' : 'Add New Route' }}
                </h3>
                <p class="text-sm text-gray-500 mt-1 font-medium">{{ isEditing ? 'Modify the details of this operational route.' : 'Define a new operational route.' }}</p>
              </div>
            </div>
            <UButton color="neutral" variant="ghost" icon="i-lucide-x" class="-my-1 rounded-full hover:bg-gray-100 transition-colors" @click="() => { isOpen = false }" />
          </div>
        </template>

        <form @submit.prevent="$emit('save')" class="space-y-6 px-2 py-4">
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
            <UFormField label="Origin City" required class="w-full">
              <UInput v-model="form.originCity" required placeholder="e.g. Jakarta" class="w-full" size="xl" icon="i-lucide-building" :ui="{ base: 'w-full' }" />
            </UFormField>
            <UFormField label="Origin Pool" required class="w-full">
              <UInput v-model="form.originPool" required placeholder="e.g. Lebak Bulus" class="w-full" size="xl" icon="i-lucide-map-pin" :ui="{ base: 'w-full' }" />
            </UFormField>
          </div>
          
          <div class="flex justify-center -my-2 text-gray-300">
            <UIcon name="i-lucide-arrow-down" class="w-6 h-6" />
          </div>
          
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
            <UFormField label="Destination City" required class="w-full">
              <UInput v-model="form.destinationCity" required placeholder="e.g. Bandung" class="w-full" size="xl" icon="i-lucide-building" :ui="{ base: 'w-full' }" />
            </UFormField>
            <UFormField label="Destination Pool" required class="w-full">
              <UInput v-model="form.destinationPool" required placeholder="e.g. Cihampelas" class="w-full" size="xl" icon="i-lucide-map-pin" :ui="{ base: 'w-full' }" />
            </UFormField>
          </div>

          <div class="flex justify-end gap-3 mt-10 pt-8 border-t border-gray-100">
            <UButton color="neutral" variant="soft" size="lg" class="px-6 rounded-xl font-semibold hover:bg-gray-100" @click="() => { isOpen = false }">Cancel</UButton>
            <UButton type="submit" color="primary" size="lg" class="px-6 rounded-xl font-semibold shadow-md hover:shadow-lg transition-all hover:-translate-y-0.5 active:translate-y-0">
              {{ isEditing ? 'Save Changes' : 'Create Route' }}
            </UButton>
          </div>
        </form>
      </UCard>
    </template>
  </UModal>
</template>

<script setup lang="ts">
const isOpen = defineModel<boolean>('open', { default: false })
const form = defineModel<any>('form', { required: true })

defineProps<{
  isEditing: boolean
}>()

defineEmits(['save'])
</script>
