user account for bob:
  user.present:
    - name: bob
    - shell: /bin/bash
    - home: /home/bob
    - groups:
      - sudo

