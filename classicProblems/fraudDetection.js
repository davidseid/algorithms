const parseLine = (line) => {

  let splitLine = line.split(';');
  let name = splitLine[0];
  let info = splitLine[1];

  let parsedLine = {
    name: name,
    type: null,
    invoices: null,
  }

  if (info === 'START') {
    parsedLine.type = 'start';
  } else {
    parsedLine.type = 'billing';
    parsedLine.invoices = info.split(',');
  }

  return parsedLine;
}

const createJob = (lineNo, latestInvoice) => {
  return {
    line: lineNo,
    minInvoice: latestInvoice,
    actualInvoice: null
  };
}

const addToJobs = (name, job, jobs) => {
  jobs[name] ? jobs[name].push(job) : jobs[name] = [job];
}

const detectShortenedJob = (invoice, name, jobs, frauds) => {
  let mostRecentJob = jobs[name].length - 1;
  let job = jobs[name][mostRecentJob];

  if (job.actualInvoice === null) {
    job.actualInvoice = invoice;

    if (job.minInvoice > job.actualInvoice) {
      frauds.push(`${job.line};${name};SHORTENED_JOB`);
    }
  } 
}

const countViolations = (sortedInvoices, name, jobs) => {
  let batchViolations = 0;

  for (let j = 0; j < sortedInvoices.length; j++) {
    let invoice = sortedInvoices[j];
    let job = jobs[name][j];

    if (invoice < job.minInvoice) {
      batchViolations++;
    }
  }

  return batchViolations;
}

const detectBatchViolations = (invoices, lineNo, name, jobs, frauds) => {

  let sortedInvoices = invoices.map(invoice => parseInt(invoice)).sort((a, b) => a - b);

  let batchViolations = countViolations(sortedInvoices, name, jobs);

  if (batchViolations > 0) {
    frauds.push(`${lineNo};${name};SUSPICIOUS_BATCH`);
  }

  if (batchViolations === sortedInvoices.length) {

    for (let k = 0; k < jobs[name].length; k++) {
      let job = jobs[name][k];
      frauds.push(`${job.line};${name};SHORTENED_JOB`);
    }
  }
}


const fraudDetection = (datafeed) => {
  let latestInvoice = 0;
  let jobs = {};
  let frauds = [];

  for (let i = 0; i < datafeed.length; i++) {
    const { name, type, invoices } = parseLine(datafeed[i]);

    if (type === 'start') {

      let job = createJob(i + 1, latestInvoice);
      addToJobs(name, job, jobs);

    } else if (type === 'billing') {

      if (invoices.length === 1) {
        
        latestInvoice = Math.max(latestInvoice, invoices[0]);
        detectShortenedJob(parseInt(invoices[0]), name, jobs, frauds);

      } else {

        detectBatchViolations(invoices, i + 1, name, jobs, frauds)
      }
    }
  }
  return frauds;
}

let testInput1 = [
  'Alice;START',
  'Bob;START',
  'Bob;1',
  'Carson;START',
  'Alice;15',
  'Carson;6',
  'David;START',
  'David;24',
  'Evil;START',
  'Evil;24',
  'Evil;START',
  'Evil;18',
  'Fiona;START',
];

let testInput2 = [
  'Tom;START',
  'Jeremy;START',
  'Dana;START',
  'Jeremy;4',
  'Dana;2',
  'James;START',
  'Leah;START',
  'James;5',
  'Nick;START',
  'Tom;1',
  'Nick;6',
  'Leah;3',
];

let testInput3 = [
  'Nick;START',
  'Jeremy;START',
  'Leah;START',
  'Nick;10',
  'Jeremy;START',
  'Jeremy;START',
  'Leah;15',
  'Jeremy;8,14,9',
];

console.log(fraudDetection(testInput1));
console.log(fraudDetection(testInput2));
console.log(fraudDetection(testInput3));