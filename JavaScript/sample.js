// own promise

const api = "https://dummyjson.com/todos";

const promise1 = new Promise((resolve, reject) => {
  setTimeout(()=>{
    reject(1)
  },5000)
});
const promise2 = new Promise((resolve, reject) => reject(2));
const promise3 = new Promise((resolve, reject) => reject(3));

Promise.any([promise1, promise2, promise3])
  .then((res) => console.log(res))
  .catch((e) => {
    console.log(e.errors)
  });
