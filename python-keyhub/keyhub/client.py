from oauthlib.oauth2 import BackendApplicationClient
from requests_oauthlib import OAuth2Session

class Group(object):
    def __init__(self, json):
        self.uuid = json['uuid']
        self.name = json['name']
        self.href = json['links'][0]['href']

class VaultRecord(object):
    def __init__(self, json):
        self.uuid = json['uuid']
        self.name = json['name']
        self.color = json['color'] if 'color' in json else ''
        self.username = json['username'] if 'username' in json else ''
        self.url = json['url'] if 'url' in json else ''
        self.password = json['additionalObjects']['secret']['password'] if 'secret' in json['additionalObjects'] and 'password' in json['additionalObjects']['secret'] else ''
        self.totp = json['additionalObjects']['secret']['totp'] if 'secret' in json['additionalObjects'] and 'totp' in json['additionalObjects']['secret'] else ''
        self.filename = json['filename'] if 'filename' in json else ''
        self.file = json['additionalObjects']['secret']['file'] if 'secret' in json['additionalObjects'] and 'file' in json['additionalObjects']['secret'] else ''

class KeyHubClient(object):
    def __init__(self, uri, client_id, client_secret):
        self._uri = uri

        client = BackendApplicationClient(client_id=client_id)
        oauth = OAuth2Session(client=client)
        token = oauth.fetch_token(token_url=uri+'/login/oauth2/token?authVault=access', client_id=client_id, client_secret=client_secret)

        session = OAuth2Session(client_id, token=token)
        session.headers.update({'Accept': 'application/vnd.topicus.keyhub+json;version=latest'})
        session.headers.update({'topicus-Vault-session': token['vaultSession']})
        self.session = session

    def info(self):
        # TODO: return info object instead of json
        response = self.session.get(self._uri+"/keyhub/rest/v1/info")
        return response.json()

    def get_group(self, group_uuid):
        response = self.session.get(self._uri + "/keyhub/rest/v1/group?uuid=" + group_uuid)
        return Group(response.json()['items'][0])

    def get_vault_records(self, group_uuid):
        # TODO: support for client vault -> group_uuid=None
        group = self.get_group(group_uuid)
        response = self.session.get(group.href + '/vault')
        records = []
        for record in response.json()['records']:
            records.append(VaultRecord(record))
        return records

    def get_vault_record(self, group_uuid, record_uuid):
        group = self.get_group(group_uuid)
        response = self.session.get(group.href + "/vault/record/uuid/" + record_uuid + "?additional=secret") 
        return VaultRecord(response.json())


def client(uri, client_id, client_secret):
    return KeyHubClient(uri=uri, client_id=client_id, client_secret=client_secret)

