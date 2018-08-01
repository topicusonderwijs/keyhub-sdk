from setuptools import setup

setup(
    name='keyhub',
    version='0.1.0',
    description='Python client for the Topicus KeyHub Vault API',
    keywords='keyhub',
    url='www.topicus-keyhub.com',
    author='Topicus KeyHub',
    author_email='keyhub@topicus.nl',
    license='Apache Software License 2.0',
    packages=['keyhub'],
    install_requires=[
        'requests_oauthlib>=0.8.0',
    ],
    classifiers=[
        'Programming Language :: Python :: 3',
    ],
    zip_safe=False
)
