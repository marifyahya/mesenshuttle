<template>
  <UModal v-model:open="isOpen">
    <template #content>
      <UCard class="ring-0 divide-y divide-gray-100 rounded-3xl shadow-2xl overflow-hidden bg-white/80 backdrop-blur-xl">
        <template #header>
          <div class="flex items-start justify-between px-2 pt-2">
            <div class="flex items-center gap-4">
              <div class="p-3.5 bg-linear-to-br from-primary-500 to-primary-600 text-white rounded-2xl shadow-lg shadow-primary-500/30 ring-1 ring-primary-600/20">
                <UIcon :name="isEditing ? 'i-lucide-edit-3' : 'i-lucide-calendar-plus'" class="w-6 h-6" />
              </div>
              <div>
                <h3 class="text-xl font-bold text-gray-900 tracking-tight">
                  {{ isEditing ? 'Edit Schedule Details' : 'Add New Schedule' }}
                </h3>
                <p class="text-sm text-gray-500 mt-1 font-medium">{{ isEditing ? 'Update the details for this departure.' : 'Fill in the details for a new departure.' }}</p>
              </div>
            </div>
            <UButton color="neutral" variant="ghost" icon="i-lucide-x" class="-my-1 rounded-full hover:bg-gray-100 transition-colors" @click="() => { isOpen = false }" />
          </div>
        </template>

        <form @submit.prevent="$emit('save')" class="space-y-6 px-2 py-4">
          <UFormField label="Route" required class="w-full">
            <USelect 
              v-model.number="form.routeId" 
              :items="store.routes.map(r => ({ label: `${r.originCity} (${r.originPool}) → ${r.destinationCity} (${r.destinationPool})`, value: r.id }))" 
              required 
              class="w-full" size="xl" icon="i-lucide-map"
            />
          </UFormField>
          
          <UFormField label="Fleet" required class="w-full">
            <USelect 
              v-model.number="form.fleetId" 
              :items="store.fleets.map(f => ({ label: `${f.name} - ${f.type} (${f.capacity} seats)`, value: f.id }))" 
              required 
              class="w-full" size="xl" icon="i-lucide-bus"
            />
          </UFormField>
          
          <div class="grid grid-cols-1 sm:grid-cols-2 gap-6">
            <UFormField label="Departure Time" required class="w-full">
              <UInput v-model="form.departureTime" type="datetime-local" required class="w-full" size="xl" icon="i-lucide-clock" />
            </UFormField>
            
            <UFormField label="Price (Rp)" required class="w-full">
              <UInput v-model.number="form.price" type="number" required min="0" class="w-full" size="xl" icon="i-lucide-wallet" />
            </UFormField>
          </div>

          <div class="flex justify-end gap-3 mt-10 pt-8 border-t border-gray-100">
            <UButton color="neutral" variant="soft" size="lg" class="px-6 rounded-xl font-semibold hover:bg-gray-100" @click="() => { isOpen = false }">Cancel</UButton>
            <UButton type="submit" color="primary" size="lg" class="px-6 rounded-xl font-semibold shadow-md hover:shadow-lg transition-all hover:-translate-y-0.5 active:translate-y-0">
              {{ isEditing ? 'Save Changes' : 'Create Schedule' }}
            </UButton>
          </div>
        </form>
      </UCard>
    </template>
  </UModal>
</template>

<script setup lang="ts">
import { useMockDataStore } from '~/stores/mockData'

const isOpen = defineModel<boolean>('open', { default: false })
const form = defineModel<any>('form', { required: true })

defineProps<{
  isEditing: boolean
}>()

defineEmits(['save'])

const store = useMockDataStore()
</script>
