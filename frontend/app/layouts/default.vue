<template>
  <div class="min-h-screen flex flex-col bg-gray-50">
    <Navbar />
    <main class="flex-grow">
      <slot />
    </main>
    <Footer />
  </div>
</template>
