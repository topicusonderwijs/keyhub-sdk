# Licensed to the Apache Software Foundation (ASF) under one or more
# contributor license agreements.  See the NOTICE file distributed with
# this work for additional information regarding copyright ownership.
# The ASF licenses this file to You under the Apache License, Version 2.0
# (the "License"); you may not use this file except in compliance with
# the License.  You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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
