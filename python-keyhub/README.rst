=============
python-keyhub
=============

Python client for the Topicus KeyHub Vault API, supported under python 3 (only tested with 3.6).

Basic usage:

.. code-block:: python

    import keyhub

    keyhub_client = keyhub.client(uri='<your KeyHub uri>', client_id='<KeyHub application id>', client_secret='<KeyHub application secret>')

    print(keyhub_client.info())

    keyhub_client.get_group('<group uuid>'

    keyhub_client.get_vault_records('<group uuid>')

    keyhub_client.get_vault_record('<group uuid>', '<vault record uuid>')

