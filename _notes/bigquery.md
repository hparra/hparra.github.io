BigQuery
========

Load data from a source:

  `bq load [DATASET].[TABLE_NAME] [PATH_TO_SOURCE] [SCHEMA]`

It's best to load large files to Google Cloud Storage first.
Its much faster overall and is a middle ground in case multiple imports are necessary.
See [Google Cloud Storage](https://github.com/hparra/hgpa/wiki/google-cloud-storage).
