<template>
  <div class="min-h-screen flex bg-gray-100">
    <aside class="w-64 bg-white border-r border-gray-200 hidden md:block">
      <div class="h-16 flex items-center px-6 border-b border-gray-200">
        <NuxtLink to="/admin/dashboard" class="text-xl font-bold text-primary-600">Admin Panel</NuxtLink>
      </div>
      <nav class="p-4 space-y-1">
        <NuxtLink to="/admin/dashboard" class="block px-3 py-2 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50 hover:text-primary-600" active-class="bg-primary-50 text-primary-600">Dashboard</NuxtLink>
        <NuxtLink to="/admin/transactions" class="block px-3 py-2 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50 hover:text-primary-600" active-class="bg-primary-50 text-primary-600">Transaction</NuxtLink>
        <NuxtLink to="/admin/schedules" class="block px-3 py-2 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50 hover:text-primary-600" active-class="bg-primary-50 text-primary-600">Schedule</NuxtLink>
        <NuxtLink to="/admin/routes" class="block px-3 py-2 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50 hover:text-primary-600" active-class="bg-primary-50 text-primary-600">Routes</NuxtLink>
        <NuxtLink to="/admin/fleets" class="block px-3 py-2 rounded-md text-sm font-medium text-gray-700 hover:bg-gray-50 hover:text-primary-600" active-class="bg-primary-50 text-primary-600">Fleet</NuxtLink>
      </nav>
    </aside>
    <div class="flex-1 flex flex-col">
      <header class="h-16 bg-white border-b border-gray-200 flex items-center px-6 justify-between md:justify-end">
        <div class="md:hidden text-lg font-bold text-primary-600">MesenShuttle Admin</div>
        <NuxtLink to="/" class="text-sm text-gray-500 hover:text-gray-700">Back to Site</NuxtLink>
      </header>
      <main class="flex-1 p-6 overflow-y-auto">
        <slot />
      </main>
    </div>
  </div>
</template>
