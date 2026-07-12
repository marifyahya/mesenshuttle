<template>
  <div class="relative min-h-[calc(100vh-130px)] flex flex-col items-center justify-center py-12 overflow-hidden">
    <!-- Animated Background Gradients -->
    <div class="absolute inset-0 bg-gray-50/50"></div>
    <div class="absolute -top-40 -right-40 w-96 h-96 bg-primary-400/30 rounded-full blur-3xl mix-blend-multiply animate-blob"></div>
    <div class="absolute top-40 -left-20 w-80 h-80 bg-indigo-400/30 rounded-full blur-3xl mix-blend-multiply animate-blob animation-delay-2000"></div>
    <div class="absolute -bottom-40 right-20 w-96 h-96 bg-sky-400/30 rounded-full blur-3xl mix-blend-multiply animate-blob animation-delay-4000"></div>

    <UContainer class="relative z-10 w-full">
      <div class="text-center mb-16">
        <h1 class="text-4xl font-extrabold tracking-tight text-gray-900 sm:text-5xl md:text-6xl animate-fade-in-up">
          <span class="block">Book your next journey</span>
          <span class="block bg-clip-text text-transparent bg-linear-to-r from-primary-600 to-indigo-600 drop-shadow-sm">with MesenShuttle</span>
        </h1>
        <p class="mt-4 max-w-md mx-auto text-base text-gray-600 sm:text-lg md:mt-6 md:text-xl md:max-w-3xl animate-fade-in-up animation-delay-300">
          Fast, reliable, and premium shuttle bookings without the hassle of registration.
        </p>
      </div>

      <UCard class="max-w-4xl mx-auto ring-1 ring-white/50 rounded-3xl shadow-2xl overflow-hidden bg-white/70 backdrop-blur-xl animate-fade-in-up animation-delay-600">
        <template #header>
          <div class="flex items-center gap-3">
            <div class="p-2 bg-primary-50 rounded-lg">
              <UIcon name="i-lucide-map-pin" class="w-5 h-5 text-primary-600" />
            </div>
            <h3 class="text-xl font-bold leading-6 text-gray-900">Search Schedules</h3>
          </div>
        </template>
        <div class="grid grid-cols-1 gap-y-6 sm:grid-cols-3 sm:gap-x-6 py-2">
          <UFormField label="Origin" name="origin" :error="formErrors.origin || undefined" class="font-medium">
            <USelect v-model="searchForm.origin" :items="groupedOrigins" placeholder="Select origin pool" size="xl" class="w-full shadow-sm" icon="i-lucide-map-pin" />
          </UFormField>
          <UFormField label="Destination" name="destination" :error="formErrors.destination || undefined" class="font-medium">
            <USelect v-model="searchForm.destination" :items="groupedDestinations" placeholder="Select destination pool" :disabled="!searchForm.origin" size="xl" class="w-full shadow-sm" icon="i-lucide-navigation" />
          </UFormField>
          <UFormField label="Date" name="date" class="font-medium">
            <UPopover class="w-full">
              <UButton color="neutral" icon="i-lucide-calendar" class="w-full justify-between text-gray-700 shadow-sm ring-1 ring-gray-200" size="xl">
                {{ formattedDate }}
              </UButton>
              <template #content>
                <UCalendar v-model="(searchForm.date as any)" :min-value="currentDate" />
              </template>
            </UPopover>
          </UFormField>
        </div>
        <template #footer>
          <div class="flex justify-end pt-2">
            <UButton @click="handleSearch" size="xl" color="primary" class="font-bold px-8 shadow-lg hover:shadow-xl transition-all hover:-translate-y-0.5 rounded-xl" trailing-icon="i-lucide-arrow-right">
              Search Shuttles
            </UButton>
          </div>
        </template>
      </UCard>
    </UContainer>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useRouter } from 'vue-router'
import { useMockDataStore } from '~/stores/mockData'
import { getLocalTimeZone, today } from '@internationalized/date'

const router = useRouter()
const store = useMockDataStore()

const groupedOrigins = computed(() => {
  const groups = new Map<string, string[]>()
  store.routes.forEach(r => {
    if (!groups.has(r.originCity)) groups.set(r.originCity, [])
    if (!groups.get(r.originCity)!.includes(r.originPool)) {
      groups.get(r.originCity)!.push(r.originPool)
    }
  })

  const result = []
  for (const [city, pools] of groups.entries()) {
    const group = [
      { label: city, type: 'label', class: 'font-bold text-lg text-primary-600' },
      ...pools.map(pool => ({ label: pool, value: pool }))
    ]
    result.push(group)
  }
  return result
})

const selectedOriginCity = computed(() => {
  if (!searchForm.value.origin) return null;
  const route = store.routes.find(r => r.originPool === searchForm.value.origin);
  return route ? route.originCity : null;
})

const selectedDestinationCity = computed(() => {
  if (!searchForm.value.destination) return null;
  const route = store.routes.find(r => r.destinationPool === searchForm.value.destination);
  return route ? route.destinationCity : null;
})

const groupedDestinations = computed(() => {
  const groups = new Map<string, string[]>()
  store.routes.forEach(r => {
    // Do not include any pools from the same city as the selected origin
    if (r.destinationCity !== selectedOriginCity.value) {
      if (!groups.has(r.destinationCity)) groups.set(r.destinationCity, [])
      if (!groups.get(r.destinationCity)!.includes(r.destinationPool)) {
        groups.get(r.destinationCity)!.push(r.destinationPool)
      }
    }
  })

  const result = []
  for (const [city, pools] of groups.entries()) {
    const group = [
      { label: city, type: 'label', class: 'font-bold text-lg text-primary-600' },
      ...pools.map(pool => ({ label: pool, value: pool }))
    ]
    result.push(group)
  }
  return result
})

// Get today's date using internationalized date
const currentDate = today(getLocalTimeZone())

const searchForm = ref({
  origin: '',
  destination: '',
  date: currentDate
})

const formErrors = ref({
  origin: '',
  destination: ''
})

const formattedDate = computed(() => {
  if (!searchForm.value.date) return 'Select date'
  const d = searchForm.value.date
  const day = d.day.toString().padStart(2, '0')
  const month = d.month.toString().padStart(2, '0')
  const year = d.year
  return `${day}-${month}-${year}`
})

// Automatically reset destination if it conflicts with the new origin city, and clear origin error
watch(() => searchForm.value.origin, () => {
  formErrors.value.origin = ''
  if (selectedOriginCity.value && selectedOriginCity.value === selectedDestinationCity.value) {
    searchForm.value.destination = ''
  }
})

// Clear destination error when a destination is selected
watch(() => searchForm.value.destination, () => {
  formErrors.value.destination = ''
})

const handleSearch = () => {
  formErrors.value = { origin: '', destination: '' }
  let hasError = false
  
  if (!searchForm.value.origin) {
    formErrors.value.origin = 'Origin is required'
    hasError = true
  }
  if (!searchForm.value.destination) {
    formErrors.value.destination = 'Destination is required'
    hasError = true
  }
  
  if (hasError) return;
  
  // Pass state via query
  router.push({
    path: '/schedules',
    query: { origin: searchForm.value.origin, destination: searchForm.value.destination, date: searchForm.value.date.toString() }
  })
}
</script>

<style scoped>
@keyframes blob {
  0% { transform: translate(0px, 0px) scale(1); }
  33% { transform: translate(30px, -50px) scale(1.1); }
  66% { transform: translate(-20px, 20px) scale(0.9); }
  100% { transform: translate(0px, 0px) scale(1); }
}
.animate-blob {
  animation: blob 7s infinite;
}
.animation-delay-2000 {
  animation-delay: 2s;
}
.animation-delay-4000 {
  animation-delay: 4s;
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translate3d(0, 20px, 0);
  }
  to {
    opacity: 1;
    transform: translate3d(0, 0, 0);
  }
}
.animate-fade-in-up {
  animation: fadeInUp 0.8s ease-out forwards;
  opacity: 0;
}
.animation-delay-300 {
  animation-delay: 0.3s;
}
.animation-delay-600 {
  animation-delay: 0.6s;
}
</style>
