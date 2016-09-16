package rhevapi

const test_api_root = `<api>
    <link href="/api/capabilities" rel="capabilities"/>
    <link href="/api/clusters" rel="clusters"/>
    <link href="/api/clusters?search={query}" rel="clusters/search"/>
    <link href="/api/datacenters" rel="datacenters"/>
    <link href="/api/datacenters?search={query}" rel="datacenters/search"/>
    <link href="/api/events" rel="events"/>
    <link href="/api/events;from={event_id}?search={query}" rel="events/search"/>
    <link href="/api/hosts" rel="hosts"/>
    <link href="/api/hosts?search={query}" rel="hosts/search"/>
    <link href="/api/networks" rel="networks"/>
    <link href="/api/networks?search={query}" rel="networks/search"/>
    <link href="/api/roles" rel="roles"/>
    <link href="/api/storagedomains" rel="storagedomains"/>
    <link href="/api/storagedomains?search={query}" rel="storagedomains/search"/>
    <link href="/api/tags" rel="tags"/>
    <link href="/api/bookmarks" rel="bookmarks"/>
    <link href="/api/icons" rel="icons"/>
    <link href="/api/templates" rel="templates"/>
    <link href="/api/templates?search={query}" rel="templates/search"/>
    <link href="/api/instancetypes" rel="instancetypes"/>
    <link href="/api/instancetypes?search={query}" rel="instancetypes/search"/>
    <link href="/api/users" rel="users"/>
    <link href="/api/users?search={query}" rel="users/search"/>
    <link href="/api/groups" rel="groups"/>
    <link href="/api/groups?search={query}" rel="groups/search"/>
    <link href="/api/domains" rel="domains"/>
    <link href="/api/vmpools" rel="vmpools"/>
    <link href="/api/vmpools?search={query}" rel="vmpools/search"/>
    <link href="/api/vms" rel="vms"/>
    <link href="/api/vms?search={query}" rel="vms/search"/>
    <link href="/api/disks" rel="disks"/>
    <link href="/api/disks?search={query}" rel="disks/search"/>
    <link href="/api/jobs" rel="jobs"/>
    <link href="/api/storageconnections" rel="storageconnections"/>
    <link href="/api/vnicprofiles" rel="vnicprofiles"/>
    <link href="/api/diskprofiles" rel="diskprofiles"/>
    <link href="/api/cpuprofiles" rel="cpuprofiles"/>
    <link href="/api/schedulingpolicyunits" rel="schedulingpolicyunits"/>
    <link href="/api/schedulingpolicies" rel="schedulingpolicies"/>
    <link href="/api/permissions" rel="permissions"/>
    <link href="/api/macpools" rel="macpools"/>
    <link href="/api/operatingsystems" rel="operatingsystems"/>
    <link href="/api/externalhostproviders" rel="externalhostproviders"/>
    <link href="/api/openstackimageproviders" rel="openstackimageproviders"/>
    <link href="/api/openstackvolumeproviders" rel="openstackvolumeproviders"/>
    <link href="/api/openstacknetworkproviders" rel="openstacknetworkproviders"/>
    <link href="/api/katelloerrata" rel="katelloerrata"/>
    <special_objects>
        <link href="/api/templates/00000000-0000-0000-0000-000000000000" rel="templates/blank"/>
        <link href="/api/tags/00000000-0000-0000-0000-000000000000" rel="tags/root"/>
    </special_objects>
    <product_info>
        <name>Red Hat Enterprise Virtualization Manager</name>
        <vendor>Red Hat</vendor>
        <version major="3" minor="6" build="8" revision="0"/>
        <full_version>3.6.8.1-0.1.el6</full_version>
    </product_info>
    <summary>
        <vms>
            <total>35</total>
            <active>9</active>
        </vms>
            <hosts>
            <total>3</total>
            <active>3</active>
        </hosts>
        <users>
            <total>12</total>
            <active>12</active>
        </users>
        <storage_domains>
            <total>4</total>
            <active>3</active>
        </storage_domains>
    </summary>
    <time></time>
</api>`

const test_api_datacenters = `<data_centers>
    <data_center id="01a45ff0-915a-11e0-8b87-5254004ac988" href="/api/datacenters/01a45ff0-915a-11e0-8b87-5254004ac988">
	<name>Default</name>
	<description>The default Data Center</description>
	<link rel="storagedomains" href="/api/datacenters/01a45ff0-915a-11e0-8b87-5254004ac988/storagedomains"/>
	<link rel="permissions" href="/api/datacenters/01a45ff0-915a-11e0-8b87-5254004ac988/permissions"/>
	<storage_type>nfs</storage_type>
	<storage_format>v1</storage_format>
	<version minor="0" major="3"/>
	<supported_versions>
	    <version minor="0" major="3"/>
	</supported_versions>
	<status>
	    <state>up</state>
	</status>
    </data_center>
</data_centers>`

const test_api_clusters = `<clusters>
    <cluster id="99408929-82cf-4dc7-a532-9d998063fa95" href="/api/clusters/99408929-82cf-4dc7-a532-9d998063fa95">
        <name>Default</name>
        <description>The default server cluster</description>
        <link rel="networks" href="/api/clusters/99408929-82cf-4dc7-a532-9d998063fa95/networks"/>
        <link rel="permissions" href="/api/clusters/99408929-82cf-4dc7-a532-9d998063fa95/permissions"/>
        <cpu id="Intel Penryn Family"/>
        <data_center id="01a45ff0-915a-11e0-8b87-5254004ac988" href="/api/datacenters/01a45ff0-915a-11e0-8b87-5254004ac988"/>
        <memory_policy>
            <overcommit percent="100"/>
            <transparent_hugepages>
                <enabled>false</enabled>
            </transparent_hugepages>
        </memory_policy>
        <scheduling_policy/>
        <version minor="0" major="3"/>
        <error_handling>
            <on_error>migrate</on_error>
        </error_handling>
    </cluster>
</clusters>`

const test_api_networks = `<networks>
    <network id="00000000-0000-0000-0000-000000000009" href="/api/networks/00000000-0000-0000-0000-000000000009">
        <name>rhevm</name>
        <description>Management Network</description>
        <data_center id="01a45ff0-915a-11e0-8b87-5254004ac988" href="/api/datacenters/01a45ff0-915a-11e0-8b87-5254004ac988"/>
        <stp>false</stp>
        <status>
            <state>operational</state>
        </status>
        <display>false</display>
    </network>
</networks>`

const test_api_hosts = `<hosts>
    <host id="0656f432-923a-11e0-ad20-5254004ac988" href="/api/hosts/0656f432-923a-11e0-ad20-5254004ac988">
        <name>hypervisor</name>
        <actions>
            <link rel="install" href="/api/hosts/0656f432-923a-11e0-ad20-5254004ac988/install"/>
            <link rel="activate" href="/api/hosts/0656f432-923a-11e0-ad20-5254004ac988/activate"/>
            <link rel="fence" href="/api/hosts/0656f432-923a-11e0-ad20-5254004ac988/fence"/>
            <link rel="deactivate" href="/api/hosts/0656f432-923a-11e0-ad20-5254004ac988/deactivate"/>
            <link rel="approve" href="/api/hosts/0656f432-923a-11e0-ad20-5254004ac988/approve"/>
            <link rel="iscsilogin" href="/api/hosts/0656f432-923a-11e0-ad20-5254004ac988/iscsilogin"/>
            <link rel="iscsidiscover"href="/api/hosts/0656f432-923a-11e0-ad20-5254004ac988/iscsidiscover"/>
            <link rel="commitnetconfig" href="/api/hosts/0656f432-923a-11e0-ad20-5254004ac988/commitnetconfig"/>
        </actions>
        <link rel="storage" href="/api/hosts/0656f432-923a-11e0-ad20-5254004ac988/storage"/>
        <link rel="nics" href="/api/hosts/0656f432-923a-11e0-ad20-5254004ac988/nics"/>
        <link rel="tags" href="/api/hosts/0656f432-923a-11e0-ad20-5254004ac988/tags"/>
        <link rel="permissions" href="/api/hosts/0656f432-923a-11e0-ad20-5254004ac988/permissions"/>
        <link rel="statistics" href="/api/hosts/0656f432-923a-11e0-ad20-5254004ac988/statistics"/>
        <address>10.64.14.110</address>
        <status>
            <state>non_operational</state>
        </status>
        <cluster id="99408929-82cf-4dc7-a532-9d998063fa95" href="/api/clusters/99408929-82cf-4dc7-a532-9d998063fa95"/>
        <port>54321</port>
        <storage_manager>true</storage_manager>
        <power_management>
            <enabled>false</enabled>
            <options/>
        </power_management>
        <ksm>
            <enabled>false</enabled>
        </ksm>
        <transparent_hugepages>
            <enabled>true</enabled>
        </transparent_hugepages>
        <iscsi>
            <initiator>iqn.1994-05.com.example:644949fe81ce</initiator>
        </iscsi>
        <cpu>
            <topology cores="2"/>
            <name>Intel(R) Core(TM)2 Duo CPU E8400 @ 3.00GHz</name>
            <speed>2993</speed>
        </cpu>
        <summary>
            <active>0</active>
            <migrating>0</migrating>
            <total>0</total>
        </summary>
    </host>
</hosts>`

const test_api_vm_createresponse = `<vm id="6efc0cfa-8495-4a96-93e5-ee490328cf48" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48">
    <name>test</name>
    <actions>
        <link rel="shutdown" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/shutdown"/>
        <link rel="start" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/start"/>
        <link rel="stop" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/stop"/>
        <link rel="suspend" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/suspend"/>
        <link rel="detach" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/detach"/>
        <link rel="export" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/export"/>
        <link rel="move" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/move"/>
        <link rel="ticket" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/ticket"/>
        <link rel="migrate" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/migrate"/>
    </actions>
    <link rel="disks" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/disks"/>
    <link rel="nics" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/nics"/>
    <link rel="cdroms" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/cdroms"/>
    <link rel="snapshots" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/snapshots"/>
    <link rel="tags" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/tags"/>
    <link rel="permissions" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/permissions"/>
    <link rel="statistics" href="/api/vms/6efc0cfa-8495-4a96-93e5-ee490328cf48/statistics"/>
    <type>server</type>
    <status>
        <state>down</state>
    </status>
    <memory>1073741824</memory>
    <cpu>
        <topology cores="1" sockets="1"/>
    </cpu>
    <os type="Unassigned">
        <boot dev="hd"/>
    </os>
    <high_availability>
        <enabled>false</enabled>
        <priority>0</priority>
    </high_availability>
    <display>
        <type>spice</type>
        <monitors>1</monitors>
    </display>
    <cluster id="99408929-82cf-4dc7-a532-9d998063fa95" href="/api/clusters/99408929-82cf-4dc7-a532-9d998063fa95"/>
    <template id="00000000-0000-0000-0000-000000000000" href="/api/templates/00000000-0000-0000-0000-000000000000"/>
    <start_time>2011-06-15T04:48:02.167Z</start_time>
    <creation_time>2011-06-15T14:48:02.078+10:00</creation_time>
    <origin>rhev</origin>
    <stateless>false</stateless>
    <placement_policy>
        <affinity>migratable</affinity>
    </placement_policy>
    <memory_policy>
        <guaranteed>536870912</guaranteed>
    </memory_policy>
</vm>`
