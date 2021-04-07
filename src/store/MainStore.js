import { reactive, computed } from 'vue'

const state = reactive({
  m3uData: [],
  m3uDataLength: 0,
})

//MUTATIONS
function PUSH_DATA(data) {
  state.m3uData = data
  state.m3uDataLength = data.length
}

// ACTIONS
async function AddM3uData(data) {
  PUSH_DATA(data)
}

// GETTERS
const GetM3uData = computed(() => state.m3uData)
const LengthM3uData = computed(() => state.m3uDataLength)

export { GetM3uData, AddM3uData, LengthM3uData }
