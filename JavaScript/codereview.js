// Function to calculate the sum of an array
function calculateSum(array) {
  let sum = 0
  for (let i = 0; i < array.length; i++) {
    sum += array[i]
  }
  return sum
}

// Function to fetch user data from an API
async function fetchUserData(userId) {
  const url = "https://jsonplaceholder.typicode.com/users/" + userId
  let response = await fetch(url)
  let data = await response.json()
  console.log("User Data:", data)
  return data
}

// Function to format a user's full name
function formatFullName(user) {
  return user.first_name + " " + user.last_name
}

// Test the functions
console.log("Sum:", calculateSum([1, 2, 3, 4]))
fetchUserData(1).then((user) => {
  console.log("Formatted Name:", formatFullName(user))
})
