<template>
  <div class="min-h-[calc(100vh-130px)] flex items-center justify-center py-12 px-4 sm:px-6 lg:px-8 relative overflow-hidden">
    <!-- Animated Background Gradients -->
    <div class="absolute inset-0 bg-gray-50/50"></div>
    <div class="absolute -top-40 -right-40 w-96 h-96 bg-primary-400/30 rounded-full blur-3xl mix-blend-multiply animate-blob"></div>
    <div class="absolute -bottom-40 -left-40 w-96 h-96 bg-indigo-400/30 rounded-full blur-3xl mix-blend-multiply animate-blob animation-delay-2000"></div>

    <UCard class="w-full max-w-md relative z-10 ring-1 ring-gray-100 rounded-3xl shadow-2xl overflow-hidden bg-white/80 backdrop-blur-xl">
      <template #header>
        <div class="text-center pt-4">
          <div class="mx-auto w-16 h-16 bg-linear-to-br from-primary-500 to-primary-600 rounded-2xl flex items-center justify-center shadow-lg shadow-primary-500/30 mb-6">
            <UIcon name="i-lucide-shield-check" class="w-8 h-8 text-white" />
          </div>
          <h2 class="text-3xl font-extrabold text-gray-900 tracking-tight">Admin Portal</h2>
          <p class="mt-2 text-sm text-gray-500">Sign in to manage operations and fleets</p>
        </div>
      </template>

      <form @submit.prevent="handleLogin" class="space-y-6 pb-4">
        <UFormField label="Email address" required>
          <UInput 
            v-model="email" 
            type="email" 
            placeholder="admin@mesenshuttle.com" 
            size="xl" 
            icon="i-lucide-mail"
            required
            class="w-full"
          />
        </UFormField>

        <UFormField label="Password" required>
          <div class="flex items-center justify-between mb-1 absolute right-0 -top-6">
            <a href="#" class="text-xs font-medium text-primary-600 hover:text-primary-500">Forgot password?</a>
          </div>
          <UInput 
            v-model="password" 
            type="password" 
            placeholder="••••••••" 
            size="xl" 
            icon="i-lucide-lock"
            required
            class="w-full"
          />
        </UFormField>

        <UButton 
          type="submit" 
          color="primary" 
          size="xl" 
          block 
          :loading="isLoading"
          class="font-bold shadow-md hover:shadow-lg transition-all hover:-translate-y-0.5 mt-8 rounded-xl"
        >
          Sign In
        </UButton>
      </form>
    </UCard>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'

definePageMeta({
  layout: 'default'
})

const router = useRouter()
const email = ref('admin@mesenshuttle.com')
const password = ref('password123')
const isLoading = ref(false)

const handleLogin = async () => {
  isLoading.value = true
  
  // Simulate API delay for premium feel
  await new Promise(resolve => setTimeout(resolve, 800))
  
  isLoading.value = false
  router.push('/admin/dashboard')
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
</style>
