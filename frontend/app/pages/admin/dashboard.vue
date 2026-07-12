<template>
  <div class="space-y-6 pb-12">
    <!-- Header -->
    <div class="flex flex-col md:flex-row md:items-center justify-between gap-4">
      <div>
        <h1 class="text-3xl font-bold text-gray-900 tracking-tight">Good Morning, Admin 👋</h1>
        <p class="text-gray-500 mt-1">Here's what's happening with MesenShuttle today, {{ currentDate }}.</p>
      </div>
      <div class="flex items-center gap-3">
        <UButton color="neutral" variant="solid" icon="i-lucide-calendar" size="md">
          Generate Report
        </UButton>
      </div>
    </div>

    <!-- Metrics Grid -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <!-- Total Revenue Card -->
      <div class="relative overflow-hidden rounded-2xl bg-linear-to-br from-indigo-500 to-purple-600 p-6 text-white shadow-lg hover:-translate-y-1 hover:shadow-xl transition-all duration-300">
        <div class="relative z-10 flex flex-col h-full justify-between">
          <div class="flex items-center justify-between">
            <span class="text-indigo-100 font-medium text-sm">Total Revenue</span>
            <div class="p-2 bg-white/20 rounded-lg backdrop-blur-sm">
              <UIcon name="i-lucide-wallet" class="w-5 h-5 text-white" />
            </div>
          </div>
          <div class="mt-4">
            <h3 class="text-3xl font-bold">Rp {{ totalRevenue.toLocaleString('id-ID') }}</h3>
            <p class="text-indigo-100 text-xs mt-1 flex items-center gap-1">
              <UIcon name="i-lucide-trending-up" class="w-3 h-3" />
              +12% from last month
            </p>
          </div>
        </div>
        <div class="absolute -right-6 -top-6 w-24 h-24 bg-white/10 rounded-full blur-2xl"></div>
      </div>

      <!-- Transactions Card -->
      <div class="relative overflow-hidden rounded-2xl bg-linear-to-br from-emerald-400 to-teal-600 p-6 text-white shadow-lg hover:-translate-y-1 hover:shadow-xl transition-all duration-300">
        <div class="relative z-10 flex flex-col h-full justify-between">
          <div class="flex items-center justify-between">
            <span class="text-emerald-100 font-medium text-sm">Transactions</span>
            <div class="p-2 bg-white/20 rounded-lg backdrop-blur-sm">
              <UIcon name="i-lucide-receipt" class="w-5 h-5 text-white" />
            </div>
          </div>
          <div class="mt-4">
            <h3 class="text-3xl font-bold">{{ store.transactions?.length || 0 }}</h3>
            <p class="text-emerald-100 text-xs mt-1">Total bookings recorded</p>
          </div>
        </div>
        <div class="absolute -right-6 -top-6 w-24 h-24 bg-white/10 rounded-full blur-2xl"></div>
      </div>

      <!-- Active Schedules Card -->
      <div class="relative overflow-hidden rounded-2xl bg-linear-to-br from-blue-500 to-sky-600 p-6 text-white shadow-lg hover:-translate-y-1 hover:shadow-xl transition-all duration-300">
        <div class="relative z-10 flex flex-col h-full justify-between">
          <div class="flex items-center justify-between">
            <span class="text-blue-100 font-medium text-sm">Active Schedules</span>
            <div class="p-2 bg-white/20 rounded-lg backdrop-blur-sm">
              <UIcon name="i-lucide-calendar-clock" class="w-5 h-5 text-white" />
            </div>
          </div>
          <div class="mt-4">
            <h3 class="text-3xl font-bold">{{ store.schedules.length }}</h3>
            <p class="text-blue-100 text-xs mt-1">Upcoming trips scheduled</p>
          </div>
        </div>
        <div class="absolute -right-6 -top-6 w-24 h-24 bg-white/10 rounded-full blur-2xl"></div>
      </div>

      <!-- Total Routes Card -->
      <div class="relative overflow-hidden rounded-2xl bg-linear-to-br from-amber-400 to-orange-500 p-6 text-white shadow-lg hover:-translate-y-1 hover:shadow-xl transition-all duration-300">
        <div class="relative z-10 flex flex-col h-full justify-between">
          <div class="flex items-center justify-between">
            <span class="text-amber-100 font-medium text-sm">Total Routes</span>
            <div class="p-2 bg-white/20 rounded-lg backdrop-blur-sm">
              <UIcon name="i-lucide-map" class="w-5 h-5 text-white" />
            </div>
          </div>
          <div class="mt-4">
            <h3 class="text-3xl font-bold">{{ store.routes.length }}</h3>
            <p class="text-amber-100 text-xs mt-1">Operational routes</p>
          </div>
        </div>
        <div class="absolute -right-6 -top-6 w-24 h-24 bg-white/10 rounded-full blur-2xl"></div>
      </div>
    </div>

    <div class="grid grid-cols-1 lg:grid-cols-3 gap-6">
      <!-- Left Column (Wider) -->
      <div class="lg:col-span-2 space-y-6">
        <!-- Recent Transactions -->
        <UCard class="ring-1 ring-gray-100 rounded-2xl shadow-sm">
          <template #header>
            <div class="flex items-center justify-between">
              <h3 class="text-lg font-bold text-gray-900">Recent Transactions</h3>
              <UButton to="/admin/transactions" color="neutral" variant="ghost" size="sm">View All</UButton>
            </div>
          </template>
          
          <div class="overflow-x-auto">
            <table class="w-full text-sm text-left">
              <thead class="text-xs text-gray-500 bg-gray-50/50">
                <tr>
                  <th class="px-4 py-3 font-medium rounded-tl-lg">Booking Code</th>
                  <th class="px-4 py-3 font-medium">Customer</th>
                  <th class="px-4 py-3 font-medium">Amount</th>
                  <th class="px-4 py-3 font-medium rounded-tr-lg">Status</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-gray-100">
                <tr v-for="txn in recentTransactions" :key="txn.id" class="hover:bg-gray-50/50 transition-colors">
                  <td class="px-4 py-3 font-medium text-gray-900">{{ txn.bookingCode }}</td>
                  <td class="px-4 py-3 text-gray-600">
                    <div>{{ txn.customerName }}</div>
                    <div class="text-xs text-gray-400">{{ txn.email }}</div>
                  </td>
                  <td class="px-4 py-3 text-gray-900 font-medium">Rp {{ txn.amount.toLocaleString('id-ID') }}</td>
                  <td class="px-4 py-3">
                    <UBadge :color="txn.status === 'PAID' ? 'success' : (txn.status === 'PENDING' ? 'warning' : 'error')" variant="subtle" size="xs">
                      {{ txn.status }}
                    </UBadge>
                  </td>
                </tr>
                <tr v-if="recentTransactions.length === 0">
                  <td colspan="4" class="px-4 py-8 text-center text-gray-500">No transactions found</td>
                </tr>
              </tbody>
            </table>
          </div>
        </UCard>
      </div>

      <!-- Right Column -->
      <div class="space-y-6">
        <!-- Quick Actions -->
        <UCard class="ring-1 ring-gray-100 rounded-2xl shadow-sm" :ui="{ body: 'p-4' }">
          <template #header>
            <h3 class="text-lg font-bold text-gray-900">Quick Actions</h3>
          </template>
          <div class="grid grid-cols-2 gap-3">
            <NuxtLink to="/admin/schedules" class="flex flex-col items-center justify-center p-4 bg-gray-50 rounded-xl hover:bg-primary-50 hover:text-primary-600 transition-colors group cursor-pointer border border-gray-100 hover:border-primary-100">
              <div class="w-10 h-10 bg-white rounded-full flex items-center justify-center shadow-sm text-gray-500 group-hover:text-primary-600 mb-2">
                <UIcon name="i-lucide-calendar-plus" class="w-5 h-5" />
              </div>
              <span class="text-xs font-medium text-gray-700 group-hover:text-primary-700">New Schedule</span>
            </NuxtLink>
            
            <NuxtLink to="/admin/routes" class="flex flex-col items-center justify-center p-4 bg-gray-50 rounded-xl hover:bg-primary-50 hover:text-primary-600 transition-colors group cursor-pointer border border-gray-100 hover:border-primary-100">
              <div class="w-10 h-10 bg-white rounded-full flex items-center justify-center shadow-sm text-gray-500 group-hover:text-primary-600 mb-2">
                <UIcon name="i-lucide-map-pin" class="w-5 h-5" />
              </div>
              <span class="text-xs font-medium text-gray-700 group-hover:text-primary-700">New Route</span>
            </NuxtLink>
            
            <NuxtLink to="/admin/fleets" class="flex flex-col items-center justify-center p-4 bg-gray-50 rounded-xl hover:bg-primary-50 hover:text-primary-600 transition-colors group cursor-pointer border border-gray-100 hover:border-primary-100">
              <div class="w-10 h-10 bg-white rounded-full flex items-center justify-center shadow-sm text-gray-500 group-hover:text-primary-600 mb-2">
                <UIcon name="i-lucide-bus" class="w-5 h-5" />
              </div>
              <span class="text-xs font-medium text-gray-700 group-hover:text-primary-700">Add Fleet</span>
            </NuxtLink>
            
            <NuxtLink to="/admin/transactions" class="flex flex-col items-center justify-center p-4 bg-gray-50 rounded-xl hover:bg-primary-50 hover:text-primary-600 transition-colors group cursor-pointer border border-gray-100 hover:border-primary-100">
              <div class="w-10 h-10 bg-white rounded-full flex items-center justify-center shadow-sm text-gray-500 group-hover:text-primary-600 mb-2">
                <UIcon name="i-lucide-receipt" class="w-5 h-5" />
              </div>
              <span class="text-xs font-medium text-gray-700 group-hover:text-primary-700">Transactions</span>
            </NuxtLink>
          </div>
        </UCard>

        <!-- Upcoming Departures -->
        <UCard class="ring-1 ring-gray-100 rounded-2xl shadow-sm">
          <template #header>
            <h3 class="text-lg font-bold text-gray-900">Upcoming Departures</h3>
          </template>
          
          <div class="space-y-4">
            <template v-for="schedule in upcomingSchedules" :key="schedule?.id || Math.random()">
              <div v-if="schedule" class="flex gap-4 items-start p-3 rounded-xl hover:bg-gray-50 transition-colors border border-transparent hover:border-gray-100">
                <div class="flex flex-col items-center justify-center min-w-[50px] bg-primary-50 text-primary-700 rounded-lg p-2 text-center">
                  <span class="text-xs font-semibold">{{ formatTimeOnly(schedule.departureTime) }}</span>
                </div>
                <div class="flex-1">
                  <h4 class="text-sm font-bold text-gray-900">{{ schedule.route?.originCity }} → {{ schedule.route?.destinationCity }}</h4>
                  <p class="text-xs text-gray-500 mt-0.5">{{ schedule.fleet?.name }} • {{ schedule.availableSeats }} seats left</p>
                </div>
              </div>
            </template>
            
            <div v-if="upcomingSchedules.length === 0" class="text-center py-4 text-sm text-gray-500">
              No upcoming schedules today.
            </div>
          </div>
        </UCard>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useMockDataStore } from '~/stores/mockData'

definePageMeta({
  layout: 'admin'
})

const store = useMockDataStore()

// Current Date formatting
const currentDate = computed(() => {
  return new Date().toLocaleDateString('en-US', { weekday: 'long', year: 'numeric', month: 'long', day: 'numeric' })
})

// Total Revenue calculation
const totalRevenue = computed(() => {
  if (!store.transactions) return 0
  return store.transactions
    .filter(t => t.status === 'PAID')
    .reduce((sum, t) => sum + t.amount, 0)
})

// Recent Transactions (last 5)
const recentTransactions = computed(() => {
  if (!store.transactions) return []
  return [...store.transactions]
    .sort((a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime())
    .slice(0, 5)
})

// Upcoming Schedules (Next 4)
const upcomingSchedules = computed(() => {
  const now = new Date().getTime()
  return store.schedules
    .map(s => store.getScheduleDetails(s.id))
    .filter(s => s && new Date(s.departureTime).getTime() > now)
    .sort((a, b) => new Date(a!.departureTime).getTime() - new Date(b!.departureTime).getTime())
    .slice(0, 4)
})

const formatTimeOnly = (dateString: string) => {
  return new Date(dateString).toLocaleTimeString('id-ID', { hour: '2-digit', minute: '2-digit' })
}
</script>
