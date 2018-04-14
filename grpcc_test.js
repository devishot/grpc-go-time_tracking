console.log('createRecord');
console.log('allRecords');
console.log('');

let record = { id: 'qwer-asdf-zxcv-hjkl', description: 'my work', amount: 120, timestamp: 1523727281 }
client.createRecord(record, pr);

client.allRecords({}, pr);
