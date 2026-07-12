<template>
  <UModal v-model:open="isOpen">
    <template #content>
      <UCard class="ring-0 divide-y divide-gray-100 rounded-3xl shadow-2xl overflow-hidden bg-white/80 backdrop-blur-xl">
        <template #header>
          <div class="flex items-start justify-between px-2 pt-2">
            <div class="flex items-center gap-4">
              <div class="p-3.5 bg-linear-to-br from-primary-500 to-primary-600 text-white rounded-2xl shadow-lg shadow-primary-500/30 ring-1 ring-primary-600/20">
                <UIcon :name="isEditing ? 'i-lucide-edit-3' : 'i-lucide-bus'" class="w-6 h-6" />
              </div>
              <div>
                <h3 class="text-xl font-bold text-gray-900 tracking-tight">
                  {{ isEditing ? 'Edit Fleet Details' : 'Register New Fleet' }}
                </h3>
                <p class="text-sm text-gray-500 mt-1 font-medium">{{ isEditing ? 'Update the settings for this vehicle.' : 'Add a new vehicle to your operational fleet.' }}</p>
              </div>
            </div>
            <UButton color="neutral" variant="ghost" icon="i-lucide-x" class="-my-1 rounded-full hover:bg-gray-100 transition-colors" @click="() => { isOpen = false }" />
          </div>
        </template>

        <form @submit.prevent="$emit('save')" class="space-y-6 px-2 py-4">
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
            <UFormField label="Fleet Name" required class="w-full">
              <UInput v-model="form.name" required placeholder="e.g. MesenShuttle 08" class="w-full" size="xl" icon="i-lucide-tag" />
            </UFormField>
            <UFormField label="Plat Number" required class="w-full">
              <UInput v-model="form.platNumber" required placeholder="e.g. B 1234 XYZ" class="w-full" size="xl" icon="i-lucide-hash" />
            </UFormField>
          </div>
          
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
            <UFormField label="Fleet Type" required class="w-full">
              <USelect v-model="form.type" :items="['Reguler', 'Premium']" required class="w-full" size="xl" icon="i-lucide-star" />
            </UFormField>
            <UFormField label="Capacity (Seats)" required class="w-full">
              <UInput v-model.number="form.capacity" type="number" required min="1" max="50" class="w-full" size="xl" icon="i-lucide-users" />
            </UFormField>
          </div>

          <div class="flex justify-end gap-3 mt-10 pt-8 border-t border-gray-100">
            <UButton color="neutral" variant="soft" size="lg" class="px-6 rounded-xl font-semibold hover:bg-gray-100" @click="() => { isOpen = false }">Cancel</UButton>
            <UButton type="submit" color="primary" size="lg" class="px-6 rounded-xl font-semibold shadow-md hover:shadow-lg transition-all hover:-translate-y-0.5 active:translate-y-0">
              {{ isEditing ? 'Save Changes' : 'Create Fleet' }}
            </UButton>
          </div>
        </form>
      </UCard>
    </template>
  </UModal>
</template>

<script setup lang="ts">
import { watch } from 'vue'

const isOpen = defineModel<boolean>('open', { default: false })
const form = defineModel<any>('form', { required: true })

const props = defineProps<{
  isEditing: boolean
}>()

defineEmits(['save'])

// Auto update capacity based on type for convenience
watch(() => form.value.type, (newType) => {
  if (!props.isEditing) {
    form.value.capacity = newType === 'Premium' ? 10 : 12
  }
})
</script>
