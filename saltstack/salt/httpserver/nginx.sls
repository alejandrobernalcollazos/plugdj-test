install nginx:
  pkg.installed:
    - name: nginx

set the nginx conf file:
  file.managed:
    - name: /etc/nginx/sites-available/default
    - source: salt://resources/nginx/default

restart nginx:
   module.run:
    - name: service.restart
    - m_name: nginx
