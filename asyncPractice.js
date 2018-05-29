// resolve after 2

const resolveAfter2Secs = () => {
  console.log('start slow promise');
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve();
      console.log('slow promise resolves')
    }, 1000);
  });
}

const resolveAfter1Sec = () => {
  console.log('start fast promise');
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve();
      console.log('fast promise resolves');
    }, 2000)
  });
}

// sequential

const sequentialStart = async () => {
  console.log('--Sequential start --');
  await resolveAfter2Secs();
  await resolveAfter1Sec();
  console.log('--End Sequential --');
}

sequentialStart();

// parallel

// concurrent

// stil serial