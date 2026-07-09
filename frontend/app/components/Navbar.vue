<template>
  <header class="bg-white border-b border-gray-200">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between h-16 items-center">
        <div class="flex-shrink-0 flex items-center">
          <NuxtLink to="/" class="text-2xl font-bold text-primary-600">
            MesenShuttle
          </NuxtLink>
        </div>
        <nav class="hidden md:flex space-x-8">
          <NuxtLink to="/" class="text-gray-700 hover:text-primary-600 px-3 py-2 text-sm font-medium">Home</NuxtLink>
          <NuxtLink to="/admin/login" class="text-gray-700 hover:text-primary-600 px-3 py-2 text-sm font-medium">Admin Portal</NuxtLink>
        </nav>
      </div>
    </div>
  </header>
</template>
