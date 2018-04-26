let assert = require('assert');
let gClient = require('grpc.client');

const protofile = '../grpc-protofiles/time_tracking/api.proto';
const client = gClient({address: 'localhost:9000'});

var createdRecordID = null;
const userID = '406ed9e1-6983-4081-a063-7dd71b3078c1';
const projectID = '9dd8c5f7-5651-4223-80af-621d7ab73539';

describe('TimeTracking', function() {
    describe('#createRecord', function() {
        it('should return stored record with new `id` and `timestamp`', function (done) {
            let record = {
                description: 'my work', amount: 30,
                user_id: userID,
                project_id: projectID};

            client
                .service('TimeTracking', protofile)
                .createRecord(record)
                .end((err, res, metadata) => {
                    if (err) {
                        done(err);
                        return;
                    }

                    assert.notEqual(res.id, '');
                    assert.notEqual(res.timestamp, '');
                    assert.equal(res.description, record.description);
                    assert.equal(res.amount, record.amount);
                    assert.equal(res.user_id, record.user_id);
                    assert.equal(res.project_id, record.project_id);

                    createdRecordID = res.id;
                    console.log(res.id);
                    done();
                })
        });
    })
});

describe('TimeTracking', function() {
    describe('#allRecords', function() {
        it('should return records for `user_id`', function (done) {
            let allRecordsRequest = { user_id: userID };

            client
                .service('TimeTracking', protofile)
                .allRecords(allRecordsRequest)
                .end((err, res, metadata) => {
                    if (err) {
                        done(err);
                        return;
                    }

                    console.log(res);

                    let all = res.records;
                    let nonEmpty = all.length >= 1;
                    let ids = all.map(x => x.id);
                    let includeJustCreated = ids.includes(createdRecordID);

                    assert.ok(nonEmpty);
                    assert.ok(includeJustCreated);


                    done();
                })
        });
    })
});

describe('TimeTracking', function() {
    describe('#allRecords', function() {
        it('should return records for `project_id`', function (done) {
            let allRecordsRequest = { project_id: projectID };

            client
                .service('TimeTracking', protofile)
                .allRecords(allRecordsRequest)
                .end((err, res, metadata) => {
                    if (err) {
                        done(err);
                        return;
                    }

                    console.log(res);

                    let all = res.records;
                    let nonEmpty = all.length >= 1;
                    let ids = all.map(x => x.id);
                    let includeJustCreated = ids.includes(createdRecordID);

                    assert.ok(nonEmpty);
                    assert.ok(includeJustCreated);


                    done();
                })
        });
    })
});

describe('TimeTracking', function() {
    describe('#deleteRecord', function() {
        it('should return deleted record with `id`', function (done) {
            let deleteRecordRequest = {id: createdRecordID};

            client
                .service('TimeTracking', protofile)
                .deleteRecord(deleteRecordRequest)
                .end((err, res, metadata) => {
                    if (err) {
                        done(err);
                        return;
                    }

                    assert.equal(res.id, deleteRecordRequest.id);
                    console.log(res.id);
                    done();
                })
        });
    })
});
