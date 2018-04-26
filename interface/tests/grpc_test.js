let assert = require('assert');
let gClient = require('grpc.client');

const protofile = '../grpc-protofiles/time_tracking/api.proto';
const client = gClient({address: 'localhost:9000'});

var createdRecordID = null;

describe('TimeTracking', function() {
    describe('#createRecord', function() {
        it('should return stored record with exist `id` and `timestamp`', function (done) {
            let record = {description: 'my work', amount: 30};

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

                    createdRecordID = res.id;
                    console.log(res.id);
                    done();
                })
        });
        //TODO: test for existance of `user_id` and `project_id`
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
