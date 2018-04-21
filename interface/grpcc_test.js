console.log('createRecord');
console.log('allRecords');
console.log('');

let record = { description: 'my work', amount: 30 }
client.createRecord(record, pr);

client.allRecords({}, pr);
